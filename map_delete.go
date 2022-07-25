package main

import "fmt"

func main() {
	mapa := map[string][]string{
		"carlos_antonio": {"video game", "read", "program"},
		"boarro_lara":    {"play with dogs", "program", "drink"},
		"luiz_eduardo":   {"nargs", "drink", "play cards"},
	}

	mapa["junior_paulo"] = []string{"test", "gin", "work"}

	delete(mapa, "carlos_antonio")

	for key, value := range mapa {
		fmt.Print(key, " like: ")
		for _, val := range value {
			fmt.Print(val, " ")
		}
		fmt.Print("\n")
	}
}
