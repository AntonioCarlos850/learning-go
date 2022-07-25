package main

import "fmt"

type person struct {
	name                string
	last_name           string
	favorite_ice_creams []string
}

func main() {
	first := person{
		name:                "Antonio",
		last_name:           "Carlos",
		favorite_ice_creams: []string{"chocolate", "the blue one"},
	}

	second := person{
		"Bruna",
		"da Silva",
		[]string{"vanilla", "cream", "chocolate"},
	}

	fmt.Printf("%v %v likes eat: ", first.name, first.last_name)
	lenght := len(first.favorite_ice_creams) - 1
	for key, value := range first.favorite_ice_creams {
		if key == lenght {
			fmt.Print(value, "\n")
			continue
		}
		fmt.Print(value, " and ")
	}

	fmt.Printf("%v %v likes eat: ", second.name, second.last_name)
	lenght = len(second.favorite_ice_creams) - 1
	for key, value := range second.favorite_ice_creams {
		if key == lenght {
			fmt.Print(value, "\n")
			continue
		}
		fmt.Print(value, " and ")
	}
}
