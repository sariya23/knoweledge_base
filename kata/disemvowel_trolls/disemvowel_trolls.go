package main

import (
	"fmt"
	"strings"
)

func Disemvowel(comment string) string {
	vowels := "oeuia"
	result := ""

	for _, v := range comment {
		if !strings.Contains(vowels, strings.ToLower(string(v))) {
			result += string(v)
		}
	}
	return result
}

func main() {
	fmt.Println(Disemvowel("XYz"))
}
