package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2, 3, 2))
	fmt.Println(sub(20, 1, 2, 3, 2))
}

func sum(x ...int) int {
	total := 0
	for _, value := range x {
		total += value
	}
	return total
}

func sub(initial_value int, x ...int) (string, int) {
	for _, value := range x {
		initial_value -= value
	}
	return "The result is", initial_value
}
