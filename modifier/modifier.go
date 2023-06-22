package modifier

import "github.com/i2y/wasabi"

func Compose(mods ...func(wasabi.Element) wasabi.Element) func(wasabi.Element) wasabi.Element {
	return func(e wasabi.Element) wasabi.Element {
		return e.Modify(mods...)
	}
}

func Attr(name, value string) func(wasabi.Element) wasabi.Element {
	return func(e wasabi.Element) wasabi.Element {
		return e.Attr(name, value)
	}
}

func Style(name, value string) func(wasabi.Element) wasabi.Element {
	return func(e wasabi.Element) wasabi.Element {
		return e.Style(name, value)
	}
}
