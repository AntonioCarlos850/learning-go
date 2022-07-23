package main

import "fmt"

const (
	_ = iota
	x
	y
	z
)

func main() {
	sum := x + y + z
	fmt.Println(x, y, z)
	fmt.Print(sum)
}