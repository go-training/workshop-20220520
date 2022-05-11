package main

import "fmt"

type StatusType int

const (
	StatusPrepare StatusType = iota
	StatusInitial
	StatusRunning
	StatusSuccess
	StatusFailure
)

func main() {
	fmt.Println(
		StatusInitial,
		StatusRunning,
		StatusSuccess,
		StatusFailure,
	) // 0, 1, 2, 3
}
