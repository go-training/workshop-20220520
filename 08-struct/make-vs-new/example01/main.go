package main

import (
	"fmt"
)

type example struct {
	foo string
	bar string
}

func main() {
	a := &example{}
	b := new(example)

	fmt.Println(a.foo)
	fmt.Println(b.foo)

	c := &example{
		foo: "foo",
		bar: "bar",
	}

	fmt.Println(c.foo)
}
