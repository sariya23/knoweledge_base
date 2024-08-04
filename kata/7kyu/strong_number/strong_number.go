package main

import "fmt"

const (
	strong    = "STRONG!!!!"
	notStrong = "Not Strong !!"
)

func factorial(n int) int {
	if n == 0 {
		return 1
	}

	result := 1

	for i := 1; i < n+1; i++ {
		result *= i
	}

	return result
}

func Strong(n int) string {
	var sumOfDigitFactorials int
	initN := n

	for n > 0 {
		if n < 10 {
			sumOfDigitFactorials += factorial(n)
		} else {
			sumOfDigitFactorials += factorial(n % 10)
		}
		n /= 10
	}

	if sumOfDigitFactorials == initN {
		return strong
	}
	return notStrong
}

func main() {
	fmt.Println(Strong(12))
}
