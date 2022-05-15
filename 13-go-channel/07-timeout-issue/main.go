package main

import (
	"fmt"
	"sync"
	"time"
)

func job1(done chan bool) {
	time.Sleep(100 * time.Millisecond)
	done <- true
}

func job2(done chan bool) {
	time.Sleep(600 * time.Millisecond)
	done <- true
}

func doJob(wg *sync.WaitGroup, f func(chan bool)) error {
	done := make(chan bool)
	defer wg.Done()
	go f(done)
	select {
	case <-done:
		fmt.Println("done")
		return nil
	case <-time.After(500 * time.Millisecond):
		return fmt.Errorf("timeout")
	}
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(4)
	go func() {
		if err := doJob(wg, job1); err != nil {
			fmt.Println(err)
		}
	}()
	go func() {
		if err := doJob(wg, job2); err != nil {
			fmt.Println(err)
		}
	}()
	go func() {
		if err := doJob(wg, job1); err != nil {
			fmt.Println(err)
		}
	}()
	go func() {
		if err := doJob(wg, job2); err != nil {
			fmt.Println(err)
		}
	}()
	wg.Wait()
}
