package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	count := 0
	num_of_func := 20

	var wg sync.WaitGroup

	wg.Add(num_of_func)

	for i := 0; i < num_of_func; i++ {
		go func() {
			v := count + 1
			runtime.Gosched()
			count = v
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(count)
}
