package main

import (
	"fmt"
	"runtime"
)

func test() chan int {
	ch := make(chan int)
	for i := 0; i < 10; i++ {
		go func(i int) {
			ch <- i
		}(i)
	}

	return ch
}

func main() {
	for i := 0; i < 4; i++ {
		test()
		fmt.Printf("goroutines: %d\n", runtime.NumGoroutine())
	}
}
