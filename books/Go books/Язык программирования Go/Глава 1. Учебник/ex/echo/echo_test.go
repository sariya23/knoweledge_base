package echo_test

import (
	"echo/echo"
	"strings"
	"testing"
)

var global string

func BenchmarkJoinViaForManyLongLetters(b *testing.B) {
	args := make([]string, 0, 10_000)
	for i := 0; i < 10_000; i++ {
		args = append(args, strings.Repeat("a", 10_000))
	}
	b.ResetTimer()

	var res string
	for i := 0; i < b.N; i++ {
		res = echo.ConcatFor(args)
	}
	global = res
}

func BenchmarkJoinViaForManyShortLetters(b *testing.B) {
	args := make([]string, 0, 10_000)
	for i := 0; i < 10_000; i++ {
		args = append(args, "a")
	}
	b.ResetTimer()

	var res string
	for i := 0; i < b.N; i++ {
		res = echo.ConcatFor(args)
	}
	global = res
}

func BenchmarkJoinViaConcatBuilderManyLongLetters(b *testing.B) {
	args := make([]string, 0, 10_000)
	for i := 0; i < 10_000; i++ {
		args = append(args, strings.Repeat("a", 10_000))
	}
	total := 0
	for i := 0; i < len(args); i++ {
		total += len(args[i])
	}
	b.ResetTimer()

	var res string
	for i := 0; i < b.N; i++ {
		res = echo.ConcatBuilder(args, total)
	}
	global = res
}

func BenchmarkJoinViaConcatBuilderManyShortLetters(b *testing.B) {
	args := make([]string, 0, 10_000)
	for i := 0; i < 10_000; i++ {
		args = append(args, "a")
	}
	total := 0
	for i := 0; i < len(args); i++ {
		total += len(args[i])
	}
	b.ResetTimer()

	var res string
	for i := 0; i < b.N; i++ {
		res = echo.ConcatBuilder(args, total)
	}
	global = res
}

func BenchmarkJoinViaAllocConcatBuilderManyShortLetters(b *testing.B) {
	args := make([]string, 0, 10_000)
	for i := 0; i < 10_000; i++ {
		args = append(args, "a")
	}
	b.ResetTimer()

	var res string
	for i := 0; i < b.N; i++ {
		res = echo.ConcatAndAllocBuilder(args)
	}
	global = res
}

func BenchmarkJoinViaAllocConcatBuilderManyLongLetters(b *testing.B) {
	args := make([]string, 0, 10_000)
	for i := 0; i < 10_000; i++ {
		args = append(args, strings.Repeat("a", 10_000))
	}
	b.ResetTimer()

	var res string
	for i := 0; i < b.N; i++ {
		res = echo.ConcatAndAllocBuilder(args)
	}
	global = res
}

func BenchmarkJoinViaJoinManyLongLetters(b *testing.B) {
	args := make([]string, 0, 10_000)
	for i := 0; i < 10_000; i++ {
		args = append(args, strings.Repeat("a", 10_000))
	}
	b.ResetTimer()

	var res string
	for i := 0; i < b.N; i++ {
		res = echo.ConcatJoin(args)
	}
	global = res
}

func BenchmarkJoinViaJoinManyShortLetters(b *testing.B) {
	args := make([]string, 0, 10_000)
	for i := 0; i < 10_000; i++ {
		args = append(args, strings.Repeat("a", 10_000))
	}
	b.ResetTimer()

	var res string
	for i := 0; i < b.N; i++ {
		res = echo.ConcatJoin(args)
	}
	global = res
}
