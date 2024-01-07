package main

import (
	"fmt"
	"math"
)

type Shapes interface {
	Area() float64
}

type Cube struct {
	X int64
}

type Circle struct {
	R float64
}

type Triangle struct {
	A, H float64
}

func (c *Cube) Area() float64 {
	return float64(c.X * c.X)
}

func (cr *Circle) Area() float64 {
	return math.Pi * cr.R * cr.R
}

func (t *Triangle) Area() float64 {
	return (t.A * t.H) / 2
}

func printInfo(shape Shapes) {
	fmt.Printf("Area of %T is %0.2f \n", shape, shape.Area())
}

func main() {
	var shape Shapes
	cube := Cube{10}
	shape = &cube
	fmt.Println(&cube)
	printInfo(shape)

	circle := Circle{
		R: 10,
	}
	shape = &circle
	fmt.Println(&circle)
	printInfo(shape)

	triangle := Triangle{A: 7, H: 12}
	shape = &triangle
	fmt.Println(&triangle)
	printInfo(shape)
}
