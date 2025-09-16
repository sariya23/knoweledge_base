package main

import "strings"

func ModifyMultiply(str string, loc, num int) string {
	if num == 0 {
		num = 1
	}
	words := strings.Split(str, " ")
	word := words[loc]

	repetedWord := make([]string, num)

	for i := 0; i < num; i++ {
		repetedWord[i] = word
	}

	return strings.Join(repetedWord, "-")
}
