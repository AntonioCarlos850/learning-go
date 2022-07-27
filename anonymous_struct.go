package main

import "fmt"

func main() {
	x := struct {
		name    string
		address []string
		phones  map[string]int
	}{
		name:    "Antonio",
		address: []string{"Rua João Holler", "Almirante Barroso"},
		phones: map[string]int{
			"home": 78953641,
			"work": 2448967315,
		},
	}

	fmt.Println(x)
}