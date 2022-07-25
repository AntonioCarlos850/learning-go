package main

import "fmt"

func main() {
	slice := []int{11, 33, 54, 22, 65, 8, 7, 9, 67, 55}

	fmt.Println("First to third: ", slice[:3])
	fmt.Println("Five until it's over: ", slice[4:])
	fmt.Println("Second to seven: ", slice[1:7])
	fmt.Println("Third to the penultimate: ", slice[2:len(slice)-1])
}
6