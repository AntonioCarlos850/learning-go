package main

import "fmt"

type square struct {
	side float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return 2 * 3.14 * c.radius
}

func (s square) area() float64 {
	return s.side * s.side
}

type figure interface {
	area() float64
}

func calc(f figure) float64 {
	return f.area()
}

func main() {
	x := square{10.0}
	y := circle{2.0}
	fmt.Println(calc(x))
	fmt.Println(calc(y))
}
