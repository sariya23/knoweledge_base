package main

import (
	"fmt"
	"strings"
)

func reverseVowels(s string) string {
	if len(s) == 1 {
		return s
	}
	i := 0
	j := len(s) - 1
	res := []byte(s)
	isIVowel, isJVowel := false, false

	for i <= j {
		fmt.Println(i, j)
		if !strings.Contains("aeiou", strings.ToLower(string(res[i]))) && !isIVowel {
			i++
		} else if strings.Contains("aeiou", strings.ToLower(string(res[i]))) {
			isIVowel = true
		}
		if !strings.Contains("aeiou", strings.ToLower(string(res[j]))) && !isJVowel {
			j--
		} else if strings.Contains("aeiou", strings.ToLower(string(res[j]))) {
			isJVowel = true
		}
		if strings.Contains("aeiou", strings.ToLower(string(res[i]))) && strings.Contains("aeiou", strings.ToLower(string(res[j]))) {
			res[i], res[j] = res[j], res[i]
			isIVowel, isJVowel = false, false
			i++
			j--
		}
	}
	return string(res)
}

func main() {
	fmt.Println(reverseVowels("race a car")) // AceCreIm
	// fmt.Println(reverseVowels("leetcode"))
}
