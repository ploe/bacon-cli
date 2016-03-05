package main

import (
	"fmt"
)

func format(c int, s string) string {
	return fmt.Sprintf("\x1b[%d;1m%s\x1b[0m", c, s)
}

var fgcode = map[string]int {
	"black":   30,
	"red":     31,
	"green":   32,	
	"yellow":  33,	
	"blue":    34,	
	"magenta": 35,	
	"cyan":    36,
} 

type fg struct {}

func (f fg) Black (s string) string {
	return format(fgcode["black"], s)
}

func (f fg) Red (s string) string {
	return format(fgcode["red"], s)
}

func (f fg) Green (s string) string {
	return format(fgcode["green"], s)
}

func (f fg) Yellow (s string) string {
	return format(fgcode["yellow"], s)
}

func (f fg) Blue (s string) string {
	return format(fgcode["blue"], s)
}

func (f fg) Magenta (s string) string {
	return format(fgcode["magenta"], s)
}

func (f fg) Cyan (s string) string {
	return format(fgcode["cyan"], s)
}

var Fg fg

func main() {
	fmt.Println(
		Fg.Red("hi my name is myke"),
		"this is normal...",
		Fg.Green("wahaha"))

	fmt.Println(
		Fg.Black("Black"),
		Fg.Red("Red"),
		Fg.Green("Green"),
		Fg.Yellow("Yellow"),
		Fg.Blue("Blue"),
		Fg.Magenta("Magenta"),
		Fg.Cyan("Cyan"),
	)
}
