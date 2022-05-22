package main

import "fmt"

func squares(count int, ch chan<- int) {
}

func main() {
	count := 5
	ch := make(chan int)

	for i := 0; i < count; i++ {
		count += <-ch
	}

	fmt.Println(count)
}
