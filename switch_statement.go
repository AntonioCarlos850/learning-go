package main

import "fmt"

func main() {
	x := "Facebook"

	switch x {
	case "Youtube":
		fmt.Print("Google app")
	case "Facebook":
		fmt.Print("Meta app")
	default:
		fmt.Print("Other")
	}
}
