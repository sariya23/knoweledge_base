package main

import "fmt"

func NbYear(p0 int, percent float64, aug int, p int) int {
	counter := 0

	for p0 < p {
		p0 = int(float64(p0) + float64(p0)*float64(percent/100) + float64(aug))
		counter += 1
	}
	return counter
}

func main() {
	fmt.Println(NbYear(1500, 5, 100, 5000))
}
