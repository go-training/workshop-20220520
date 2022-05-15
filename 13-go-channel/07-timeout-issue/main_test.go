package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestTimeoutJob02(t *testing.T) {
	wg := &sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			if err := doJob(wg, job2); err != nil {
				fmt.Println(err)
			}
		}()
	}
	wg.Wait()
}

func TestTimeoutJob01(t *testing.T) {
	wg := &sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			if err := doJob(wg, job1); err != nil {
				fmt.Println(err)
			}
		}()
	}
	wg.Wait()
}
