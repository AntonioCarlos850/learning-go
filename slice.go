package main

import "fmt"

func main() {
	slice := []int{5, 8, 77, 65, 24}
	fmt.Println(slice)

	slice = append(slice, 98)

	for key, value := range slice {
		fmt.Println("The key is", key, "and the value is", value)
	}
}
