package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	wg.Add(2)

	go first()
	go second()

	wg.Wait()
}

func first() {
	fmt.Println("First go routine")
	wg.Done()
}

func second() {
	fmt.Println("Second go routine")
	wg.Done()
}
