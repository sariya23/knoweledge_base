package main

import (
	"strconv"
	"strings"
)

func NbDig(n int, d int) int {
	counter := 0

	for i := 0; i < n+1; i++ {
		square := i * i
		stringSquareRep := strconv.Itoa(square)
		stringDigitRep := strconv.Itoa(d)
		counter += strings.Count(stringSquareRep, stringDigitRep)
	}
	return counter
}
