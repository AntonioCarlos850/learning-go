package main

import "fmt"

func main() {
	slice := [][]string{
		[]string{
			"Name", "Surname", "Hobby",
		},
		[]string{
			"João", "Carlos", "fishing",
		},
		[]string{
			"Matheus", "Felipe", "reading",
		},
		[]string{
			"Ana", "Silva", "watch tv",
		},
	}

	for _, value := range slice {
		for _, val := range value {
			fmt.Print(val, "\t\t")
		}
		fmt.Print("\n")
	}
}
