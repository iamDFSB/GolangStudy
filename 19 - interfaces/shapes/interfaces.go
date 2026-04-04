package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

func writeArea(f shape) {
	fmt.Printf("A área do shape é %0.2f \n", f.area())
}

type rectangle struct {
	height float64
	width float64
}

func (r rectangle) area() float64{
	return r.width * r.height
}

type circle struct {
	radius float64
}

func (c circle) area() float64{
	return math.Pi * math.Pow(c.radius, 2)
}


func main(){
	r := rectangle{10, 15}
	writeArea(r)

	c := circle{10}
	writeArea(c)
}	