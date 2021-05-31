package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type square struct {
	X float64
}
type circle struct {
	R float64
}

func (s square) Area() float64 {
	return s.X * s.X
}
func (s square) Perimeter() float64 {
	return 4 * s.X
}
func (c circle) Area() float64 {
	return c.R * c.R * math.Pi
}
func (c circle) Perimeter() float64 {
	return 2 * c.R * math.Pi

}

func Calculate(x Shape) {
	_, ok := x.(circle)
	if ok {
		fmt.Println("Is a Circle")
	}
	v, ok := x.(square)
	if ok {
		fmt.Println("Is a Square", v)
	}
	fmt.Println(x.Area())
	fmt.Println(x.Perimeter())
}

func main() {
	x := square{X: 10}
	fmt.Println("Perimeter: ", x.Perimeter())
	Calculate(x)
	y := circle{R: 5}
	Calculate(y)
}
