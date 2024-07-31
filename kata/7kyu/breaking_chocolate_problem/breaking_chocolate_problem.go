package main

func BreakChocolate(n, m int) int {
	if n*m-1 < 0 {
		return 0
	}
	return n*m - 1
}
