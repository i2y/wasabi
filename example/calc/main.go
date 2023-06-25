package main

import (
	"embed"
	"fmt"
	"strconv"
	"strings"

	w "github.com/i2y/wasabi"
	a "github.com/i2y/wasabi/modifier/attr"
)

//go:embed assets
var assets embed.FS

func main() {
	app := w.NewDesktopApp("Calc", 380, 486, w.Assets(assets))
	app.Run(calc)
}

type Calc struct {
	display           *w.State[string]
	operator          string
	result, operand   float64
	waitingForOperand bool
}

func NewCalc() *Calc {
	return &Calc{
		display: w.NewState("0"),
	}
}

func (c *Calc) NumberClicked(number int) {
	if c.waitingForOperand {
		c.display.Set(strconv.Itoa(number))
		c.waitingForOperand = false
	} else {
		if c.display.Get() == "0" {
			c.display.Set(strconv.Itoa(number))
		} else {
			c.display.Set(c.display.Get() + strconv.Itoa(number))
		}
	}
}

func (c *Calc) OperatorClicked(operator string) {
	if !c.waitingForOperand {
		c.operand, _ = strconv.ParseFloat(c.display.Get(), 64)
		switch c.operator {
		case "+":
			c.result += c.operand
		case "-":
			c.result -= c.operand
		case "×":
			c.result *= c.operand
		case "÷":
			if c.operand != 0 {
				c.result /= c.operand
			} else {
				c.display.Set("Error")
				return
			}
		default:
			c.result = c.operand
		}
		c.display.Set(fmt.Sprintf("%g", c.result))
	}
	c.operator = operator
	c.waitingForOperand = true
}

func (c *Calc) EqualsClicked() {
	c.OperatorClicked("=")
	c.operator = ""
}

func (c *Calc) DecimalPointClicked() {
	if c.waitingForOperand {
		c.display.Set("0")
		c.waitingForOperand = false
	}
	if !strings.Contains(c.display.Get(), ".") {
		c.display.Set(c.display.Get() + ".")
	}
}

func (c *Calc) ClearClicked() {
	c.display.Set("0")
	c.operator = ""
	c.result = 0
	c.operand = 0
	c.waitingForOperand = false
}

func calc(f *w.Factory) w.Element {
	c := NewCalc()
	return f.Div(a.Class("bg-white shadow-xl p-5"))(
		f.Reactive(c.display, func() w.Element {
			return f.Input(
				a.Class("w-full p-3 text-right text-white text-2xl mb-5"),
				a.Readonly(true),
				a.Value(c.display.Get()),
			)()
		}),
		f.Div(a.Class("grid grid-cols-4 gap-2"))(
			clear(f, c, 3), op(f, c, "÷", 1),
			num(f, c, 7, 1), num(f, c, 8, 1), num(f, c, 9, 1), op(f, c, "×", 1),
			num(f, c, 4, 1), num(f, c, 5, 1), num(f, c, 6, 1), op(f, c, "-", 1),
			num(f, c, 1, 1), num(f, c, 2, 1), num(f, c, 3, 1), op(f, c, "+", 1),
			num(f, c, 0, 2), dot(f, c, 1), eq(f, c, 1),
		),
	)
}

func num(f *w.Factory, calc *Calc, num, colspan int) w.Element {
	return f.Button(
		a.Class("btn btn-primary col-span-"+strconv.Itoa(colspan)),
		a.OnClick(func() {
			calc.NumberClicked(num)
		}),
	)(f.Text(strconv.Itoa(num)))
}

func op(f *w.Factory, calc *Calc, op string, colspan int) w.Element {
	return f.Button(
		a.Class("btn btn-accent col-span-"+strconv.Itoa(colspan)),
		a.OnClick(func() {
			calc.OperatorClicked(op)
		}),
	)(f.Text(op))
}

func eq(f *w.Factory, calc *Calc, colspan int) w.Element {
	return f.Button(
		a.Class("btn btn-secondary col-span-"+strconv.Itoa(colspan)),
		a.OnClick(func() {
			calc.EqualsClicked()
		}),
	)(f.Text("="))
}

func dot(f *w.Factory, calc *Calc, colspan int) w.Element {
	return f.Button(
		a.Class("btn btn-primary col-span-"+strconv.Itoa(colspan)),
		a.OnClick(func() {
			calc.DecimalPointClicked()
		}),
	)(f.Text("."))
}

func clear(f *w.Factory, calc *Calc, colspan int) w.Element {
	return f.Button(
		a.Class("btn btn-accent col-span-"+strconv.Itoa(colspan)),
		a.OnClick(func() {
			calc.ClearClicked()
		}),
	)(f.Text("AC"))
}
