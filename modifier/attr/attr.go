package attr

import (
	"strconv"

	"github.com/i2y/wasabi"
	"github.com/i2y/wasabi/unit"
)

func Class(names ...string) func(wasabi.Element) wasabi.Element {
	return func(e wasabi.Element) wasabi.Element {
		return e.Class(names...)
	}
}

func OnClick(handler func()) func(wasabi.Element) wasabi.Element {
	return func(e wasabi.Element) wasabi.Element {
		return e.OnClick(handler)
	}
}

func OnChange(handler func(wasabi.Event)) func(wasabi.Element) wasabi.Element {
	return func(e wasabi.Element) wasabi.Element {
		return e.OnChange(handler)
	}
}

func Width(value int, u unit.Unit) func(wasabi.Element) wasabi.Element {
	return func(e wasabi.Element) wasabi.Element {
		return e.Attr("width", strconv.Itoa(value)+string(u))
	}
}

func Height(value int, u unit.Unit) func(wasabi.Element) wasabi.Element {
	return func(e wasabi.Element) wasabi.Element {
		return e.Attr("height", strconv.Itoa(value)+string(u))
	}
}

func Href(value string) func(wasabi.Element) wasabi.Element {
	return func(e wasabi.Element) wasabi.Element {
		return e.Attr("href", value)
	}
}

func Src(value string) func(wasabi.Element) wasabi.Element {
	return func(e wasabi.Element) wasabi.Element {
		return e.Attr("src", value)
	}
}

func Alt(value string) func(wasabi.Element) wasabi.Element {
	return func(e wasabi.Element) wasabi.Element {
		return e.Attr("alt", value)
	}
}

func Value(value string) func(wasabi.Element) wasabi.Element {
	return func(e wasabi.Element) wasabi.Element {
		return e.Attr("value", value)
	}
}

func Disabled(value bool) func(wasabi.Element) wasabi.Element {
	return func(e wasabi.Element) wasabi.Element {
		return e.Attr("disabled", strconv.FormatBool(value))
	}
}

func Maxlength(value int) func(wasabi.Element) wasabi.Element {
	return func(e wasabi.Element) wasabi.Element {
		return e.Attr("maxlength", strconv.Itoa(value))
	}
}

func Placeholder(value string) func(wasabi.Element) wasabi.Element {
	return func(e wasabi.Element) wasabi.Element {
		return e.Attr("placeholder", value)
	}
}

func Readonly(value bool) func(wasabi.Element) wasabi.Element {
	return func(e wasabi.Element) wasabi.Element {
		return e.Attr("readonly", strconv.FormatBool(value))
	}
}

func Required(value bool) func(wasabi.Element) wasabi.Element {
	return func(e wasabi.Element) wasabi.Element {
		return e.Attr("required", strconv.FormatBool(value))
	}
}

func Type(value string) func(wasabi.Element) wasabi.Element {
	return func(e wasabi.Element) wasabi.Element {
		return e.Attr("type", value)
	}
}
