package main

func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	factorial := 1
	for i := 1; i < n+1; i++ {
		factorial *= i
	}

	return factorial
}
