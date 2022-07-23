package main

import "fmt"

const x int = 1

var KB int = x << 10

func main() {
	fmt.Printf("%b", KB)
}