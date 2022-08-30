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
	var mt sync.Mutex

	for i := 0; i < num_of_func; i++ {
		go func() {
			mt.Lock()

			v := count + 1
			runtime.Gosched()
			count = v

			wg.Done()
			mt.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println(count)
}
