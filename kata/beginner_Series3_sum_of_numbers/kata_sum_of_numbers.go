package main

import "fmt"

func GetSum(a, b int) int {
	start := a
	stop := b

	if b < a {
		start = b
		stop = a
	}
	fmt.Println(float64(start+stop) / 2.0)
	return int(float64(start+stop)/2.0) * (stop - start + 1)
}

func main() {
	fmt.Println(GetSum(0, 1))
}
