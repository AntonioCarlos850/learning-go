package main

import "fmt"

func main() {
	x := -51

	if x > 0 && x < 50 {
		fmt.Print("Between 0 and 50")
	} else if x >= 50 {
		fmt.Print("Bigger or equal than 50")
	} else {
		fmt.Print("Negative")
	}
}
