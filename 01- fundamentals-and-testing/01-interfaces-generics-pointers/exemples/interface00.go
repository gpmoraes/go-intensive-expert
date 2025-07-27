package main

import "fmt"

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func printArea(s Shape) {
	fmt.Printf("Area: %.2f\n", s.Area())
}

func main() {
	c := Circle{Radius: 10}
	r := Rectangle{Width: 5, Height: 6}

	printArea(c)
	printArea(r)
}
