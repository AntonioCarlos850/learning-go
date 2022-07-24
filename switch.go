package main

import "fmt"

func main() {
	x := "Antonio"

	switch {
	case x == "Antonio":
		fmt.Print("Me")
	case x == "Fiona":
		fmt.Print("My dog")
	default:
		fmt.Print("I don't who is")
	}
}