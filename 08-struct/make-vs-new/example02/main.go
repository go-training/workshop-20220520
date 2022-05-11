package main

import "fmt"

func test01(s int) {
	data := make([]int, 0, s)

	for i := 0; i < s; i++ {
		data = append(data, i)
	}
}

func test02(s int) {
	data := make([]int, 0)

	for i := 0; i < s; i++ {
		data = append(data, i)
	}
}

func test03(s int) {
	data := make([]int, s, s)

	for i := 0; i < s; i++ {
		data[i] = i
	}
}

func test04(s int) {
	data := make([]int, s)

	for i := 0; i < s; i++ {
		data = append(data, i)
	}
}

func main() {
	fmt.Println(make([]int, 5))
	fmt.Println(make([]int, 0, 5))
}
