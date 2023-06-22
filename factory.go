package wasabi

type Factory struct {
	app App
}

func newFactory(app App) *Factory {
	return &Factory{app: app}
}

func (f *Factory) Tag(t string) func(...func(Element) Element) func(...Element) Element {
	return func(mods ...func(Element) Element) func(...Element) Element {
		return func(chidlren ...Element) Element {
			container := containerElement(t)
			for _, child := range chidlren {
				if child == nil {
					continue
				}
				child.SetApp(f.app)
				container.Add(child)
			}
			container.SetApp(f.app)
			return container.Modify(mods...)
		}
	}
}

func (f *Factory) Reactive(state eventPublisher, body func() Element) *stateful {
	var comp *stateful
	comp = newStateful(state, func() Element {
		c := containerElement("div")
		c.id = comp.id
		c.SetApp(f.app)
		c.Add(body())
		return c
	})
	comp.SetApp(f.app)
	return comp
}

func Stateful[T eventPublisher](f *Factory, state T, body func(s T) Element) *stateful {
	var comp *stateful
	comp = newStateful(state, func() Element {
		c := containerElement("div")
		c.id = comp.id
		c.SetApp(f.app)
		c.Add(body(state))
		return c
	})
	comp.SetApp(f.app)
	return comp
}

func NewModifier(mods ...func(Element) Element) func(Element) Element {
	return func(e Element) Element {
		return e.Modify(mods...)
	}
}

func (f *Factory) A(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("a")(mods...)
}

func (f *Factory) Abbr(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("abbr")(mods...)
}

func (f *Factory) Address(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("address")(mods...)
}

func (f *Factory) Article(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("article")(mods...)
}

func (f *Factory) Aside(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("aside")(mods...)
}

func (f *Factory) Audio(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("audio")(mods...)
}

func (f *Factory) B(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("b")(mods...)
}

func (f *Factory) Bdi(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("bdi")(mods...)
}

func (f *Factory) Bdo(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("bdo")(mods...)
}

func (f *Factory) Blockquote(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("blockquote")(mods...)
}

func (f *Factory) Body(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("body")(mods...)
}

func (f *Factory) Button(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("button")(mods...)
}

func (f *Factory) Canvas(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("canvas")(mods...)
}

func (f *Factory) Caption(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("caption")(mods...)
}

func (f *Factory) Cite(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("cite")(mods...)
}

func (f *Factory) Code(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("code")(mods...)
}

func (f *Factory) Colgroup(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("colgroup")(mods...)
}

func (f *Factory) Datalist(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("datalist")(mods...)
}

func (f *Factory) Dd(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("dd")(mods...)
}

func (f *Factory) Del(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("del")(mods...)
}

func (f *Factory) Details(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("details")(mods...)
}

func (f *Factory) Dfn(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("dfn")(mods...)
}

func (f *Factory) Div(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("div")(mods...)
}

func (f *Factory) Dl(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("dl")(mods...)
}

func (f *Factory) Dt(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("dt")(mods...)
}

func (f *Factory) Em(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("em")(mods...)
}

func (f *Factory) Fieldset(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("fieldset")(mods...)
}

func (f *Factory) Figcaption(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("figcaption")(mods...)
}

func (f *Factory) Figure(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("figure")(mods...)
}

func (f *Factory) Footer(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("footer")(mods...)
}

func (f *Factory) Form(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("form")(mods...)
}

func (f *Factory) H1(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("h1")(mods...)
}

func (f *Factory) H2(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("h2")(mods...)
}

func (f *Factory) H3(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("h3")(mods...)
}

func (f *Factory) H4(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("h4")(mods...)
}

func (f *Factory) H5(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("h5")(mods...)
}

func (f *Factory) H6(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("h6")(mods...)
}

func (f *Factory) Head(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("head")(mods...)
}

func (f *Factory) Header(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("header")(mods...)
}

func (f *Factory) Hgroup(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("hgroup")(mods...)
}

func (f *Factory) Html(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("html")(mods...)
}

func (f *Factory) I(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("i")(mods...)
}

func (f *Factory) Input(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("input")(mods...)
}

func (f *Factory) Iframe(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("iframe")(mods...)
}

func (f *Factory) Ins(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("ins")(mods...)
}

func (f *Factory) Kbd(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("kbd")(mods...)
}

func (f *Factory) Label(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("label")(mods...)
}

func (f *Factory) Legend(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("legend")(mods...)
}

func (f *Factory) Li(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("li")(mods...)
}

func (f *Factory) Main(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("main")(mods...)
}

func (f *Factory) MapElement(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("map")(mods...)
}

func (f *Factory) Mark(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("mark")(mods...)
}

func (f *Factory) Menu(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("menu")(mods...)
}

func (f *Factory) Meter(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("meter")(mods...)
}

func (f *Factory) Nav(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("nav")(mods...)
}

func (f *Factory) Noscript(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("noscript")(mods...)
}

func (f *Factory) Object(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("object")(mods...)
}

func (f *Factory) Ol(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("ol")(mods...)
}

func (f *Factory) Optgroup(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("optgroup")(mods...)
}

func (f *Factory) Option(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("option")(mods...)
}

func (f *Factory) Output(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("output")(mods...)
}

func (f *Factory) P(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("p")(mods...)
}

func (f *Factory) Pre(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("pre")(mods...)
}

func (f *Factory) Progress(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("progress")(mods...)
}

func (f *Factory) Q(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("q")(mods...)
}

func (f *Factory) Rp(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("rp")(mods...)
}

func (f *Factory) Rt(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("rt")(mods...)
}

func (f *Factory) Ruby(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("ruby")(mods...)
}

func (f *Factory) S(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("s")(mods...)
}

func (f *Factory) Samp(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("samp")(mods...)
}

func (f *Factory) Script(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("script")(mods...)
}

func (f *Factory) Section(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("section")(mods...)
}

func (f *Factory) Select(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("select")(mods...)
}

func (f *Factory) Small(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("small")(mods...)
}

func (f *Factory) Span(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("span")(mods...)
}

func (f *Factory) Strong(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("strong")(mods...)
}

func (f *Factory) Style(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("style")(mods...)
}

func (f *Factory) Sub(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("sub")(mods...)
}

func (f *Factory) Summary(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("summary")(mods...)
}

func (f *Factory) Sup(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("sup")(mods...)
}

func (f *Factory) Svg(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("svg")(mods...)
}

func (f *Factory) Table(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("table")(mods...)
}

func (f *Factory) Tbody(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("tbody")(mods...)
}

func (f *Factory) Td(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("td")(mods...)
}

func (f *Factory) Template(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("template")(mods...)
}

func (f *Factory) Textarea(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("textarea")(mods...)
}

func (f *Factory) Tfoot(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("tfoot")(mods...)
}

func (f *Factory) Th(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("th")(mods...)
}

func (f *Factory) Thead(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("thead")(mods...)
}

func (f *Factory) Time(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("time")(mods...)
}

func (f *Factory) Tr(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("tr")(mods...)
}

func (f *Factory) U(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("u")(mods...)
}

func (f *Factory) Ul(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("ul")(mods...)
}

func (f *Factory) VarElement(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("var")(mods...)
}

func (f *Factory) Video(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("video")(mods...)
}

func (f *Factory) Wbr(mods ...func(Element) Element) func(...Element) Element {
	return f.Tag("wbr")(mods...)
}

func (f *Factory) Text(s string) Element {
	t := text(s)
	t.SetApp(f.app)
	return t
}

func If(cond bool, body func() Element) Element {
	if cond {
		return body()
	}
	return nil
}

func IfElse(cond bool, body func() Element, elseBody func() Element) Element {
	if cond {
		return body()
	}
	return elseBody()
}

func ForEach[T any](items []T, m func(item T) Element) []Element {
	els := make([]Element, len(items))
	for index, item := range items {
		els[index] = m(item)
	}
	return els
}
