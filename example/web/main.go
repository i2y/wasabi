package main

import (
	"embed"
	"net/http"
	"strconv"

	w "github.com/i2y/wasabi"
	a "github.com/i2y/wasabi/modifier/attr"
)

//go:embed counter/assets
var assets embed.FS

func main() {
	handler := w.NewHTTPHandler("counter", "/counter", counter, w.Assets(assets))
	http.ListenAndServe(":8080", handler)
}

func counter(f *w.Factory) w.Element {
	count := w.NewState(0)
	return f.Div(a.Class("flex items-center justify-center w-screen h-screen"))(
		f.Button(
			a.Class("btn text-4xl w-24 h-24 items-center justify-center"),
			a.OnClick(func() { count.Set(count.Get() + 1) }),
		)(
			f.Text("+"),
		),
		f.Reactive(count, func() w.Element {
			return f.Div(a.Class("text-slate-400 text-4xl w-24 h-24 flex items-center justify-center"))(
				f.Text(strconv.Itoa(count.Get())),
			)
		}),
		f.Button(
			a.Class("btn text-4xl w-24 h-24 items-center justify-center"),
			a.OnClick(func() { count.Set(count.Get() - 1) }),
		)(
			f.Text("-"),
		),
	)
}
