package main

import "fmt"

func main() {
	outChan := make(chan string)
	errChan := make(chan error)

	for i := 0; i < 100; i++ {
		go func(num int) {
		}(i)
	}

	for {
		select {
		case out := <-outChan:
			fmt.Println(out)
		case err := <-errChan:
			fmt.Println(err)
		}
	}
}
