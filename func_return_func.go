package main

import "fmt"

func main() {
	y := half()
	z := y(7)
	fmt.Println(z)
}

func half() func(x int) float64 {
	return func(x int) float64 {
		return float64(x) / 2
	}
}
