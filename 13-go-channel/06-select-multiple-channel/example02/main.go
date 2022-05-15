package main

import (
	"log"
	"time"
)

func main() {
	ch := make(chan int)
	go func(ch chan int) {
		for {
			if v, ok := <-ch; ok {
				log.Printf("val:%d\n", v)
			}
		}
	}(ch)

	tick := time.NewTicker(1 * time.Second)
	for i := 0; i < 30; i++ {
		select {
		// how to make sure all number send to ch channel?
		case ch <- i:
		case <-tick.C:
			log.Printf("%d: case <-tick.C\n", i)
		}

		time.Sleep(200 * time.Millisecond)
	}
	close(ch)
	tick.Stop()
}
