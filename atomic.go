package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var count int32
	num_of_func := 20

	var wg sync.WaitGroup

	wg.Add(num_of_func)

	for i := 0; i < num_of_func; i++ {
		go func() {
			atomic.AddInt32(&count, 1)
			v := atomic.LoadInt32(&count)
			fmt.Println(v)
			wg.Done()
		}()
	}

	wg.Wait()
}
