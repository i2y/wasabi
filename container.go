package wasabi

type container struct {
	base
	children []Element
}

func newContainer(children ...Element) container {
	return container{
		base:     newBase(),
		children: children,
	}
}

func (c *container) Attach(a App) {
	c.base.Attach(a)

	for _, c := range c.children {
		c.Attach(a)
	}
}

func (c *container) Detach() {
	c.base.Detach()

	for _, child := range c.children {
		child.Detach()
	}
}

func (c *container) Children() []Element {
	return c.children
}

func (c *container) Add(w Element) {
	c.children = append(c.children, w)
	c.Update()
}

func (c *container) Remove(el Element) {
	new := make([]Element, len(c.children))
	var i uint
	for _, child := range c.children {
		if child != el {
			new[i] = child
			i++
		}
	}
	c.children = new[:i]
	c.Update()
}

func (c *container) View() string {
	var ret string
	for _, w := range c.children {
		ret += w.View()
	}
	return ret
}

type genericContainerElement struct {
	container
	tag string
}

func containerElement(tag string, children ...Element) *genericContainerElement {
	el := &genericContainerElement{
		container: newContainer(children...),
		tag:       tag,
	}
	el.base.SetElement(el)
	return el
}

func (g *genericContainerElement) View() string {
	return g.genHTML(g.tag, g.container.View())
}
