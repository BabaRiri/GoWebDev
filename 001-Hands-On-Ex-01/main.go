package main

type square struct {
	side int
}

type circle struct {
	radius float64
}

func (s square) area() int {
	return s.side * 2
}
