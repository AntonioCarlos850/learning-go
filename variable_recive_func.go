package main

import "fmt"

func main() {
	x := func(y float64) float64 {
		return y / 10
	}

	z := x(78)
	fmt.Println(z)
}
