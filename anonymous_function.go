package main

import "fmt"

func main() {
	x := func(x int) int {
		return x * 10
	}(3)

	fmt.Println(x)
}
