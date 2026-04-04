package shape_methods

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

func WriteArea(f Shape) {
	fmt.Printf("A área do shape é %0.2f \n", f.Area())
}

type Rectangle struct {
	Height float64
	Width float64
}

func (r Rectangle) Area() float64{
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64{
	return math.Pi * math.Pow(c.Radius, 2)
}
