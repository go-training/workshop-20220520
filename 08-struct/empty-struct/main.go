package main

import "fmt"

func main() {
	// sample 01
	set := map[string]struct{}{}

	for _, item := range []string{"KeyA", "KeyA", "KeyB", "KeyC"} {
		set[item] = struct{}{}
	}

	if _, ok := set["KeyA"]; ok {
		fmt.Println("A exists") // A exists
	}

	// sample 02
	ch := make(chan struct{})
	go func() {
		<-ch
	}()
	ch <- struct{}{}

	foo := Light{}
	foo.On()
	foo.Off()
}

type Light struct{}

func (l Light) On() {
	println("On")
}

func (l Light) Off() {
	println("Off")
}
