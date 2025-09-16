package main

import (
	"fmt"
	"strings"
)

func reverseVowels(s string) string {
	if len(s) == 1 {
		return s
	}
	sb := strings.Split(s, "")
	vowels := "aeiuoAEIOU"
	i, j := 0, len(sb)-1

	for i < j {
		for i < j && !strings.Contains(vowels, string(sb[i])) {
			i++
		}
		for i < j && !strings.Contains(vowels, string(sb[j])) {
			j--
		}
		sb[i], sb[j] = sb[j], sb[i]
		i++
		j--
	}

	return strings.Join(sb, "")
}

func main() {
	fmt.Println(reverseVowels("IceCreAm")) // AceCreIm
	fmt.Println(reverseVowels("leetcode"))
}
