package main

import (
	"testing"
)

func BenchmarkFoo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		foo(10)
	}
}

func BenchmarkBar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bar(10)
	}
}
