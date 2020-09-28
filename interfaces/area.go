package main

import "fmt"

type shape interface {
	area() float64
}

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

func main() {
	s := square{
		sideLength: 4.0,
	}

	t := triangle{
		height: 12.0,
		base:   4.0,
	}
	printArea(t)
	printArea(s)
}
func (s square) area() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) area() float64 {
	return t.base * t.height
}

func printArea(s shape) {
	fmt.Println(s.area())
}
