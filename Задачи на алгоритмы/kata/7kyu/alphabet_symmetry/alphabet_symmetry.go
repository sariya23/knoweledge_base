package main

import (
	"strings"
)

func allToLower(s string) string {
	var result string

	for _, v := range s {
		result += strings.ToLower(string(v))
	}

	return result
}

func getLettersAmountLikeAlphabetPosition(s string) int {
	var counter int

	for i, letter := range s {
		if letter%97 == rune(i) {
			counter++
		}
	}
	return counter
}

func solve(slice []string) []int {
	result := make([]int, len(slice))

	for i, word := range slice {
		result[i] = getLettersAmountLikeAlphabetPosition(allToLower(string(word)))
	}

	return result
}
