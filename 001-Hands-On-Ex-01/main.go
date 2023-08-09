package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.side * 2
}

func (c circle) area() float64 {
	return math.Pi * (c.radius * c.radius)
}

type shape interface {
	area() float64
}

func printArea(sh shape) {
	fmt.Printf("Area: %f\n", sh.area())
}

func main() {
	s1 := square{
		side: 5.0,
	}

	s2 := circle{
		radius: 3.2,
	}

	printArea(s1)
	printArea(s2)
}
