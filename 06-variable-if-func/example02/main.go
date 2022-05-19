package main

import (
	"log"
)

type student struct {
	Name string
	Age  int
}

func main() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
		{Name: "apple", Age: 21},
		{Name: "zoo", Age: 20},
	}

	for _, stu := range stus {
		m[stu.Name] = &stu
	}

	for k, v := range m {
		log.Println(k, "=>", v)
	}
}
