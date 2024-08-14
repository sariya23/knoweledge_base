package main

import "fmt"

func Collatz(n int) int {
	counter := 1
	currentN := n

	for currentN != 1 {
		if currentN%2 == 0 {
			counter++
			currentN /= 2
		} else {
			counter++
			currentN = currentN*3 + 1
		}
	}

	return counter
}

func main() {
	fmt.Println(Collatz(20))
}
