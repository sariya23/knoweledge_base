package main

import (
	"strings"
)

func solve(str string) string {
	upperCaseLetters := 0
	lowerCaseLatters := 0

	for _, s := range str {
		if strings.ToUpper(string(s)) == string(s) {
			upperCaseLetters++
		} else {
			lowerCaseLatters++
		}
	}

	if upperCaseLetters > lowerCaseLatters {
		return strings.ToUpper(str)
	}

	return strings.ToLower(str)
}
