package main

import (
	"fmt"
	"sync"
)

var total = 0

func main() {
	count := 10
	wg := &sync.WaitGroup{}
	wg.Add(count)

	for i := 0; i < count; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 10; i++ {
				total += 1
			}
		}()
	}

	wg.Wait()
	fmt.Println(total)
}
