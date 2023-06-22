package wasabi

type stateful struct {
	base
	builder func() Element
	state   eventPublisher
	view    Element
}

func newStateful(state eventPublisher, builder func() Element) *stateful {
	c := &stateful{
		base:    newBase(),
		builder: builder,
		state:   state,
	}
	state.AddListener(c)
	c.base.SetElement(c)
	return c
}

func (c *stateful) Attach(a App) {
	c.base.Attach(a)
	if c.view != nil {
		c.view.Attach(a)
	}
}

func (c *stateful) Detach() {
	c.base.Detach()
	if c.view != nil {
		c.view.Detach()
	}
}

func (c *stateful) View() string {
	// return c.genHTML("div", c.build().View())
	content := c.build()
	content.MergeStyles(c.styles)
	content.MergeAttrs(c.attrs)
	return content.View()
}

func (c *stateful) build() Element {
	if c.view != nil {
		c.view.Detach()
		c.view = c.builder()
		c.view.Attach(c.app)
		return c.view
	} else {
		c.view = c.builder()
		return c.view
	}
}
