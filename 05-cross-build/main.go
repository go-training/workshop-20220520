package main

import (
	"fmt"

	"example04/hello"
	"example04/hello2"
	"example04/hello3"
)

func main() {
	fmt.Println(hello.Hello())
	fmt.Println(hello2.Hello())
	fmt.Println(hello3.Hello())
}
