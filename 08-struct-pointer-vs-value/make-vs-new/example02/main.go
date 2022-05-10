package main

func foo(s int) {
	data := make([]int, s)

	for i := 0; i < s; i++ {
		data = append(data, i)
	}
}

func bar(s int) {
	data := make([]int, 0)

	for i := 0; i < s; i++ {
		data = append(data, i)
	}
}
