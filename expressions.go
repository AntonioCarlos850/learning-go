package main

import "fmt"

func main() {
	equals := (8 == 4)
	lower_or_equal := (8 <= 4)
	bigger_or_equal := (8 >= 4)
	bigger := (8 > 4)
	lower := (8 < 4)

	fmt.Println(equals)
	fmt.Println(lower_or_equal)
	fmt.Println(bigger_or_equal)
	fmt.Println(bigger)
	fmt.Println(lower)
}