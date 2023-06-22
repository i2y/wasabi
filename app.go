package wasabi

import (
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"net"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
	"time"

	"github.com/asaskevich/EventBus"
	"golang.org/x/text/language"
	"nhooyr.io/websocket"
)

type App interface {
	postUpdate(w Element)
	addEventHandler(el Element, event string, handler func(ev Event))
	cancel()
}

type DesktopApp struct {
	name     string
	server   *http.Server
	listener net.Listener
	width    int
	height   int
	mods     []appOptionModifier
}

func NewDesktopApp(name string, width, height int, mods ...appOptionModifier) *DesktopApp {
	return &DesktopApp{
		name:   name,
		width:  width,
		height: height,
		mods:   mods,
	}
}

func (app *DesktopApp) Name() string {
	return app.name
}

type AppOpts struct {
	lang   language.Tag
	assets embed.FS
}

type appOptionModifier func(*AppOpts)

func AppLang(lang language.Tag) appOptionModifier {
	return func(opts *AppOpts) {
		opts.lang = lang
	}
}

func Assets(assets embed.FS) appOptionModifier {
	return func(opts *AppOpts) {
		opts.assets = assets
	}
}

func NewHTTPHandler(
	appName string,
	rootpath string,
	viewBuilder func(f *Factory) Element,
	mods ...appOptionModifier,
) http.Handler {
	return newHTTPHandler(appName, rootpath, func(ac *appCore) Element {
		return viewBuilder(newFactory(ac))
	}, mods...)
}

func (app *DesktopApp) Run(viewBuilder func(*Factory) Element) {
	app.run(NewHTTPHandler(app.Name(), "/", viewBuilder, app.mods...))
}

type appCore struct {
	view Element
	msgs chan string
	bus  EventBus.Bus
}

func newAppCore(viewBuilder func(app *appCore) Element) *appCore {
	app := &appCore{
		msgs: make(chan string),
		bus:  EventBus.New(),
	}
	app.view = viewBuilder(app)
	return app
}

func (ac *appCore) postUpdate(w Element) {
	ac.msgs <- fmt.Sprintf(
		`<turbo-stream action="replace" target="wasabi-%d"><template>%s</template></turbo-stream>`,
		w.ID(),
		w.View(),
	)
}

func (ac *appCore) addEventHandler(el Element, event string, handler func(ev Event)) {
	go func() {
		ac.bus.Subscribe(
			fmt.Sprintf("%d.%s", el.ID(), event),
			func(rev *rawEvent) {
				ev := &actualEvent{
					target:    el,
					eventName: rev.EventName,
					props:     rev.Props,
				}
				ev.target = el
				handler(ev)
			},
		)
	}()
}

const EOF = ""

func (ac *appCore) cancel() {
	ac.msgs <- EOF
}

func newHTTPHandler(
	appName string,
	rootpath string,
	viewBuilder func(*appCore) Element,
	mods ...appOptionModifier,
) http.Handler {
	appOpts := &AppOpts{
		lang:   language.English,
		assets: embed.FS{},
	}

	for _, mod := range mods {
		mod(appOpts)
	}

	appCoreFactory := func() *appCore {
		return newAppCore(viewBuilder)
	}
	mux := http.NewServeMux()

	tmpl, err := template.New("index").Parse(indexTmpl)
	if err != nil {
		panic(err)
	}

	MountAssets("", appOpts.assets, true)
	assetHandler := http.FileServer(assetsFS)
	if rootpath[len(rootpath)-1] != '/' {
		rootpath += "/"
	}

	mux.Handle(rootpath+"assets/", assetHandler)
	mux.HandleFunc(rootpath, func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.Execute(w, &templateParams{
			Name:         appName,
			Lang:         appOpts.lang,
			HeadElements: headElements,
			Events:       targetEvents,
		})
		if err != nil {
			log.Printf("failed to execute the index template: %v", err)
		}
	})

	mux.HandleFunc(rootpath+"ws", func(w http.ResponseWriter, r *http.Request) {
		c, err := websocket.Accept(w, r, nil)
		defer func() {
			if err != nil {
				c.Close(websocket.StatusInternalError, "failed to accept websocket connection")
			}
		}()
		if err != nil {
			return
		}

		err = handleWebSocket(r.Context(), c, appCoreFactory)
		if err != nil {
			return
		}
	})
	mux.HandleFunc(rootpath+"turbo.es2017-umd.js", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(turbo))
	})
	return mux
}

func handleWebSocket(ctx context.Context, ws *websocket.Conn, appCoreFactory func() *appCore) error {
	ac := appCoreFactory()
	ac.view.Attach(ac)

	defer ws.Close(websocket.StatusNormalClosure, "")

	page := fmt.Sprintf(
		`<turbo-stream action="replace" target="wasabi-page"><template>%s</template></turbo-stream>`,
		ac.view.View(),
	)

	err := ws.Write(ctx, websocket.MessageText, []byte(page))
	closeStatus := websocket.CloseStatus(err)
	if closeStatus == websocket.StatusNoStatusRcvd || closeStatus == websocket.StatusNormalClosure {
		return nil
	}

	if err != nil {
		return err
	}

	go func() {
		defer ac.cancel()
		for {
			var b []byte
			_, b, err := ws.Read(ctx)
			if err != nil {
				return
			}

			var ev rawEvent
			err = json.Unmarshal(b, &ev)
			if err != nil {
				return
			}

			ac.bus.Publish(ev.ID(), &ev)
		}
	}()

	for {
		msg := <-ac.msgs
		if msg == EOF {
			return nil
		}
		err = ws.Write(ctx, websocket.MessageText, []byte(msg))
		if err != nil {
			return err
		}
	}
}

func (app *DesktopApp) serve() {
	app.server.Serve(app.listener)
}

func (app *DesktopApp) port() int {
	return app.listener.Addr().(*net.TCPAddr).Port
}

func (app *DesktopApp) run(handler http.Handler) {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}
	app.listener = listener
	app.server = &http.Server{
		Addr:    listener.Addr().String(),
		Handler: handler,
	}

	go app.serve()
	webview := detectWebview()
	if webview == nil {
		log.Fatal("any supported webview not found")
		return
	}
	webview.open(app.name, app.port(), app.width, app.height)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := app.server.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}

type elementID int

func (id elementID) String() string {
	return fmt.Sprintf("wasabi-%d", id)
}

var eid elementID

var idMutex sync.Mutex

func newElementID() elementID {
	idMutex.Lock()
	defer idMutex.Unlock()
	eid++
	return eid
}

type rawEvent struct {
	Target    string                 `json:"target"`
	EventName string                 `json:"event"`
	Props     map[string]interface{} `json:"props"`
}

func (e *rawEvent) ID() string {
	return fmt.Sprintf("%s.%s", e.Target, e.EventName)
}

type Event interface {
	ID() string
	Target() Element
	EventName() string
	Props() map[string]interface{}
	Value() string
}

type actualEvent struct {
	target    Element
	eventName string
	props     map[string]interface{}
}

func (e *actualEvent) ID() string {
	return fmt.Sprintf("%s.%s", e.target.ID(), e.eventName)
}

func (e *actualEvent) Target() Element {
	return e.target
}

func (e *actualEvent) EventName() string {
	return e.eventName
}

func (e *actualEvent) Props() map[string]interface{} {
	return e.props
}

func (e *actualEvent) Value() string {
	return e.props["value"].(string)
}

//go:embed index.html
var indexTmpl string

//go:embed turbo.es2017-umd.js
var turbo string

type templateParams struct {
	Name         string
	Lang         language.Tag
	HeadElements []string
	Events       []TargetEvent
	Element      string
	Port         int
}

var headElements []string

func SetHeadElements(s ...string) {
	headElements = s
}

func addHeadElements(s string) {
	headElements = append(headElements, s)
}

type TargetEvent struct {
	Name     string
	PropName string
}

var targetEvents []TargetEvent

func SetTargetEvents(events []TargetEvent) {
	targetEvents = events
}

var assetsFS http.FileSystem

func MountAssets(mountDir string, assets embed.FS, headElementsFlag bool) {
	assetsLFS, err := fs.Sub(assets, ".")
	if err != nil {
		panic(err)
	}

	assetsFS = http.FS(assetsLFS)

	if !headElementsFlag {
		return
	}

	var walkRootDir string
	if mountDir == "" {
		walkRootDir = "."
	} else {
		walkRootDir = mountDir
	}

	err = fs.WalkDir(assetsLFS, walkRootDir, func(path string, d fs.DirEntry, err error) error {
		relatedPath := path[len(mountDir):]
		if relatedPath == "" || d.IsDir() {
			return nil
		}

		var filePath string
		if relatedPath[0] == '/' {
			filePath = relatedPath[1:]
		} else {
			filePath = relatedPath
		}

		fmt.Println(filePath)
		ext := filepath.Ext(filePath)
		switch ext {
		case ".css":
			addHeadElements(fmt.Sprintf(`<link rel="stylesheet" href="%s">`, filePath))
		case ".js":
			addHeadElements(fmt.Sprintf(`<script src="%s"></script>`, filePath))
		}

		return nil
	})

	if err != nil {
		panic(err)
	}
}
