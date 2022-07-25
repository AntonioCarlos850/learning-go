package main

import "fmt"

func main() {
	slice := []int{11, 33, 54, 22, 65, 8, 7, 9, 67, 55}

	fmt.Println("Index\t\tValue")
	for index, value := range slice {
		fmt.Println(index, "\t\t", value)
	}

	fmt.Printf("\nType of slice: %T", slice)
}
