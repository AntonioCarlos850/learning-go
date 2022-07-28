package main

import "fmt"

func main() {
	defer fmt.Println("Second")
	fmt.Println("First")
}
