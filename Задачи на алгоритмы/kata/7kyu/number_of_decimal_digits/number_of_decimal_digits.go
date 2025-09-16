package main

import "fmt"

func Digits(n uint64) int {
	counter := 1

	for n/10 > 0 {
		counter++
		n /= 10
	}

	return counter
}

func main() {
	fmt.Println(Digits(1234))
}
