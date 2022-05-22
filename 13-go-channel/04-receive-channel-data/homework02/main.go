package main

import (
	"sync"
)

func main() {
	str := []byte("foobar")
	ch := make(chan byte, len(str))
	_ = make(chan struct{})
	wg := &sync.WaitGroup{}
	wg.Add(2)

	for i := 0; i < len(str); i++ {
		ch <- str[i]
	}

	go func() {
	}()

	go func() {
	}()

	wg.Wait()
}
