package main

import (
	"testing"
)

var count = 1000

func Benchmark01(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		test01(count)
	}
}

func Benchmark02(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		test02(count)
	}
}

func Benchmark03(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		test03(count)
	}
}

func Benchmark04(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		test04(count)
	}
}
