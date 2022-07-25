package main

import "fmt"

func main() {
	array := [5]int{11, 24, 38, 49, 55}

	for key, value := range array {
		fmt.Printf("The index %v has value %v\n", key, value)
	}
}
