package wasabi

import (
	"fmt"
	"strings"

	"github.com/asaskevich/EventBus"
)

type Element interface {
	ID() elementID
	View() string
	Attach(app App)
	Detach()
	Update()
	Class(classes ...string) Element
	Attr(name, value string) Element
	GetAttr(name string) string
	MergeAttrs(update map[string]string) Element
	Style(name, value string) Element
	GetStyle(name string) string
	MergeStyles(update map[string]string) Element
	OnClick(handler func()) Element
	OnChange(handler func(ev Event)) Element
	Modify(mods ...func(e Element) Element) Element
	SetApp(app App)
}

type base struct {
	id       elementID
	attached bool
	app      App
	element  Element
	classes  []string
	attrs    map[string]string
	styles   map[string]string
}

func newBase() base {
	return base{
		id:       newElementID(),
		attached: false,
		attrs:    map[string]string{},
		styles:   map[string]string{},
	}
}

func (base *base) SetApp(app App) {
	base.app = app
}

func (b *base) ID() elementID {
	return b.id
}

func (b *base) View() string {
	return ""
}

func (b *base) genHTML(tag string, body string) string {
	return fmt.Sprintf(
		`<%s id="%s" class="%s" %s style="%s">%s</%s>`,
		tag,
		b.ID(),
		strings.Join(b.classes, " "),
		b.getAttrsStr(),
		b.getStyleStr(),
		body,
		tag,
	)
}

func (b *base) Attach(a App) {
	b.attached = true
	b.app = a
}

func (b *base) Detach() {
	b.attached = false
}

func (b *base) Class(classes ...string) Element {
	b.classes = classes
	return b.element
}

func (b *base) Attr(name, value string) Element {
	b.attrs[name] = value
	return b.element
}

func (b *base) GetAttr(name string) string {
	return b.attrs[name]
}

func (b *base) MergeAttrs(update map[string]string) Element {
	merge(b.attrs, update)
	return b.element
}

func (b *base) Styles(styles map[string]string) Element {
	b.styles = styles
	return b.element
}

func merge(a, b map[string]string) {
	for k, v := range b {
		a[k] = v
	}
}

func (b *base) MergeStyles(update map[string]string) Element {
	merge(b.styles, update)
	return b.element
}

func (b *base) Style(name, value string) Element {
	b.styles[name] = value
	return b.element
}

func (b *base) GetStyle(name string) string {
	return b.styles[name]
}

func (b *base) getAttrsStr() string {
	ret := ""
	for name, value := range b.attrs {
		ret += fmt.Sprintf("%s=\"%s\" ", name, value)
	}
	return ret
}

func (b *base) getStyleStr() string {
	ret := ""
	for name, value := range b.styles {
		ret += fmt.Sprintf("%s: %s;", name, value)
	}
	return ret
}

func (b *base) SetElement(w Element) {
	b.element = w
}

func (b *base) Update() {
	if b.attached {
		b.app.postUpdate(b.element)
	}
}

func (b *base) OnClick(handler func()) Element {
	b.app.addEventHandler(b.element, "click", func(_ Event) { handler() })
	// TODO on detach, delete the handler
	return b.element
}

func (b *base) OnChange(handler func(ev Event)) Element {
	b.app.addEventHandler(b.element, "change", handler)
	return b.element
}

func (b *base) Modify(mods ...func(Element) Element) Element {
	for _, mod := range mods {
		b.element = mod(b.element)
	}
	return b.element
}

type eventPublisher interface {
	AddListener(w Element)
	Notify()
}

type model struct {
	bus EventBus.Bus
}

func newModel() model {
	return model{
		bus: EventBus.New(),
	}
}

func (m *model) AddListener(b Element) {
	m.bus.Subscribe("update", b.Update)
}

func (m *model) Notify() {
	m.bus.Publish("update")
}

func (b *base) Build() Element {
	return b.element
}

func init() {
	SetTargetEvents([]TargetEvent{
		{Name: "click", PropName: ""},
		{Name: "change", PropName: "value"}},
	)
}
