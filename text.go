package wasabi

import "html"

type textElement struct {
	base
	content string
}

func text(s string) *textElement {
	el := &textElement{
		base:    newBase(),
		content: s,
	}
	el.base.SetElement(el)
	return el
}

func (t *textElement) View() string {
	return html.EscapeString(t.content)
}
