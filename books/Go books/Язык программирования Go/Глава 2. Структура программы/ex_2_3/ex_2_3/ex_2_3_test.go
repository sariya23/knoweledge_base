package pop_test

import (
	pop "main/ex_2_3"
	"testing"
)

var global int

func BenchmarkPopCountViaFor(b *testing.B) {
	var pc [256]byte

	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}

	b.ResetTimer()
	var v int
	for i := 0; i < b.N; i++ {
		v = pop.PopCountViaFor(pc, uint64(i))
	}
	global = v
}

func BenchmarkPopCountViaSum(b *testing.B) {
	var pc [256]byte

	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}

	b.ResetTimer()
	var v int
	for i := 0; i < b.N; i++ {
		v = pop.PopCountViaSum(pc, uint64(i))
	}
	global = v
}
