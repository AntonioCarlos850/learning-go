package main

import "fmt"

func main() {
	x := half(func(x int) int {
		return x * x
	}(5))
	fmt.Println(x)
}

func half(x int) float64 {
	return float64(x) / 2
}
