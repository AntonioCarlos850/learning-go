package main

import "fmt"

func main() {
	y := addOne()
	y()
	y()
	z := y()

	b := addOne()
	b()
	t := b()
	fmt.Println(z, t)
}

func addOne() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
