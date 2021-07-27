package main

import "math"

type Circle struct {
	Radius float32
}

func (c Circle) Area() float32 {
	return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
	Width  float32
	Height float32
}

func (r Rectangle) Area() float32 {
	return r.Width * r.Height
}

type ShareWithArea interface {
	Area() float32
}

func PrintArea(sa ShareWithArea) {
	println("Area = ", sa.Area())
}

func main() {
	c := Circle{Radius: 3}
	PrintArea(c)

	r := Rectangle{Width: 10, Height: 12}
	PrintArea(r)
}
