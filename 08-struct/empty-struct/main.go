package main

import "fmt"

type set map[string]struct{}

func (s set) Add(k string) {
	s[k] = struct{}{}
}

func (s set) Remove(k string) {
	delete(s, k)
}

func (s set) Has(k string) bool {
	_, ok := s[k]
	return ok
}

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

	monitor := Monitor{}
	monitor.On()
	monitor.Off()
}

type Light struct{}

func (l Light) On() {
	println("On")
}

func (l Light) Off() {
	println("Off")
}

type Monitor struct{}

func (l Monitor) On() {
	println("On")
}

func (l Monitor) Off() {
	println("Off")
}
