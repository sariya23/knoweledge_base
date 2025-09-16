package main

import (
	"fmt"
)

func PrevMultOfThree(n int) interface{} {
	for {
		if n <= 2 {
			return nil
		}
		if n%3 == 0 {
			return n
		}
		n /= 10
	}
}

func PrevMultOfThreeRecursuve(n int) interface{} {
	if n <= 2 {
		return nil
	}
	if n%3 == 0 {
		return n
	}
	return PrevMultOfThreeRecursuve(n / 10)
}

func main() {
	// fmt.Println(PrevMultOfThree(1234))
	fmt.Println(PrevMultOfThreeRecursuve(830800))
}
