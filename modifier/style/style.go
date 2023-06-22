package style

import (
	"strconv"

	"github.com/i2y/wasabi"
	"github.com/i2y/wasabi/unit"
)

func Padding(value int, u unit.Unit) func(wasabi.Element) wasabi.Element {
	return func(e wasabi.Element) wasabi.Element {
		return e.Style("padding", strconv.Itoa(value)+string(u))
	}
}

func Margin(value int, u unit.Unit) func(wasabi.Element) wasabi.Element {
	return func(e wasabi.Element) wasabi.Element {
		return e.Style("margin", strconv.Itoa(value)+string(u))
	}
}

func BackgroundColor(value string) func(wasabi.Element) wasabi.Element {
	return func(e wasabi.Element) wasabi.Element {
		return e.Style("background-color", value)
	}
}

func Color(value string) func(wasabi.Element) wasabi.Element {
	return func(e wasabi.Element) wasabi.Element {
		return e.Style("color", value)
	}
}

func FontSize(value int, u unit.Unit) func(wasabi.Element) wasabi.Element {
	return func(e wasabi.Element) wasabi.Element {
		return e.Style("font-size", strconv.Itoa(value)+string(u))
	}
}

func LineHeight(value int, u unit.Unit) func(wasabi.Element) wasabi.Element {
	return func(e wasabi.Element) wasabi.Element {
		return e.Style("line-height", strconv.Itoa(value)+string(u))
	}
}

type BorderStyle string

const (
	Solid  BorderStyle = "solid"
	Dashed BorderStyle = "dashed"
	Dotted BorderStyle = "dotted"
	Double BorderStyle = "double"
	None   BorderStyle = "none"
)

func Border(width int, u unit.Unit, style BorderStyle, color string) func(wasabi.Element) wasabi.Element {
	return func(e wasabi.Element) wasabi.Element {
		borderValue := strconv.Itoa(width) + string(u) + " " + string(style) + " " + color
		return e.Style("border", borderValue)
	}
}

func BorderWidth(value int, u unit.Unit) func(wasabi.Element) wasabi.Element {
	return func(e wasabi.Element) wasabi.Element {
		return e.Style("border-width", strconv.Itoa(value)+string(u))
	}
}

func BorderRadius(value int, u unit.Unit) func(wasabi.Element) wasabi.Element {
	return func(e wasabi.Element) wasabi.Element {
		return e.Style("border-radius", strconv.Itoa(value)+string(u))
	}
}

func Width(value int, u unit.Unit) func(wasabi.Element) wasabi.Element {
	return func(e wasabi.Element) wasabi.Element {
		return e.Style("width", strconv.Itoa(value)+string(u))
	}
}

func Height(value int, u unit.Unit) func(wasabi.Element) wasabi.Element {
	return func(e wasabi.Element) wasabi.Element {
		return e.Style("height", strconv.Itoa(value)+string(u))
	}
}

func Top(value int, u unit.Unit) func(wasabi.Element) wasabi.Element {
	return func(e wasabi.Element) wasabi.Element {
		return e.Style("top", strconv.Itoa(value)+string(u))
	}
}

func Right(value int, u unit.Unit) func(wasabi.Element) wasabi.Element {
	return func(e wasabi.Element) wasabi.Element {
		return e.Style("right", strconv.Itoa(value)+string(u))
	}
}

func Bottom(value int, u unit.Unit) func(wasabi.Element) wasabi.Element {
	return func(e wasabi.Element) wasabi.Element {
		return e.Style("bottom", strconv.Itoa(value)+string(u))
	}
}

func Left(value int, u unit.Unit) func(wasabi.Element) wasabi.Element {
	return func(e wasabi.Element) wasabi.Element {
		return e.Style("left", strconv.Itoa(value)+string(u))
	}
}

func LetterSpacing(value int, u unit.Unit) func(wasabi.Element) wasabi.Element {
	return func(e wasabi.Element) wasabi.Element {
		return e.Style("letter-spacing", strconv.Itoa(value)+string(u))
	}
}

func WordSpacing(value int, u unit.Unit) func(wasabi.Element) wasabi.Element {
	return func(e wasabi.Element) wasabi.Element {
		return e.Style("word-spacing", strconv.Itoa(value)+string(u))
	}
}

type PositionStyle string

const (
	Static   PositionStyle = "static"
	Relative PositionStyle = "relative"
	Absolute PositionStyle = "absolute"
	Fixed    PositionStyle = "fixed"
	Sticky   PositionStyle = "sticky"
)

func Position(position PositionStyle) func(wasabi.Element) wasabi.Element {
	return func(e wasabi.Element) wasabi.Element {
		return e.Style("position", string(position))
	}
}

type DisplayStyle string

const (
	Block       DisplayStyle = "block"
	Inline      DisplayStyle = "inline"
	FlexStyle   DisplayStyle = "flex"
	Grid        DisplayStyle = "grid"
	InlineBlock DisplayStyle = "inline-block"
	// ... and so on
)

func Display(display DisplayStyle) func(wasabi.Element) wasabi.Element {
	return func(e wasabi.Element) wasabi.Element {
		return e.Style("display", string(display))
	}
}

type OverflowStyle string

const (
	Visible OverflowStyle = "visible"
	Hidden  OverflowStyle = "hidden"
	Scroll  OverflowStyle = "scroll"
	Auto    OverflowStyle = "auto"
)

func Overflow(overflow OverflowStyle) func(wasabi.Element) wasabi.Element {
	return func(e wasabi.Element) wasabi.Element {
		return e.Style("overflow", string(overflow))
	}
}

type CursorStyle string

const (
	Default    CursorStyle = "default"
	Pointer    CursorStyle = "pointer"
	Text       CursorStyle = "text"
	Wait       CursorStyle = "wait"
	Crosshair  CursorStyle = "crosshair"
	NotAllowed CursorStyle = "not-allowed"
)

func Cursor(cursor CursorStyle) func(wasabi.Element) wasabi.Element {
	return func(e wasabi.Element) wasabi.Element {
		return e.Style("cursor", string(cursor))
	}
}

func Flex(grow int, shrink int, basis string) func(wasabi.Element) wasabi.Element {
	return func(e wasabi.Element) wasabi.Element {
		flexValue := strconv.Itoa(grow) + " " + strconv.Itoa(shrink) + " " + basis
		return e.Style("flex", flexValue)
	}
}
