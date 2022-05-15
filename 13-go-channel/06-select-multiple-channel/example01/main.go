package main

import (
	"fmt"
	"time"
)

type logEntry struct {
	message string
	time    time.Time
}

func main() {
	logCh := make(chan logEntry, 100)
	doneCh := make(chan struct{})
	go logger(logCh, doneCh)
	logCh <- logEntry{"App Start", time.Now()}
	logCh <- logEntry{"App End", time.Now()}
	time.Sleep(100 * time.Millisecond)
	doneCh <- struct{}{}
}

func logger(logCh <-chan logEntry, doneCh <-chan struct{}) {
Loop:
	for {
		select {
		case entry := <-logCh:
			fmt.Printf("%s %v\n", entry.time.Format(time.RFC3339), entry.message)
		case <-doneCh:
			fmt.Println("break the select loop")
			break Loop
		}
	}
	fmt.Println("exit the logger")
}
