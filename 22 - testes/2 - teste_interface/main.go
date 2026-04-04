package main

import (
	shape "shapes/shape_methods"
)

func main(){
	c := shape.Circle{10}
	shape.WriteArea(c)
}