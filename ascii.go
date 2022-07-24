package main

import "fmt"

func main() {
	for i := 33; i <= 122; i++ {
		fmt.Printf("deciaml: %d - Unicode: %U - Value: %v\n", i, i, string(i))
	}
}
