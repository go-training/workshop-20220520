package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func request(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	i := 0
	for {
		if _, err := http.Get(url); err == nil {
			break
		}
		i++
		if i >= 3 {
			break
		}
		time.Sleep(time.Second)
	}
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go request(fmt.Sprintf("https://192.168.1.1:8082/%d", i), &wg)
	}
	wg.Wait()
}
