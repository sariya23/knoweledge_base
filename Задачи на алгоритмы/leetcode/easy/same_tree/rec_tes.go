package main

import "testing"

var hole int

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s1, s2 := asyncSum()
		hole = s1 + s2
	}
}
