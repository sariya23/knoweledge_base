package main

import (
	"fmt"
	"strings"
)

func mergeAlternately(word1 string, word2 string) string {
	sb := strings.Builder{}
	sb.Grow(len(word1) + len(word2))

	i, j := 0, 0

	for {
		if i == len(word1) || j == len(word2) {
			break
		}
		_, _ = sb.WriteString(string(word1[i]))
		_, _ = sb.WriteString(string(word2[j]))
		i++
		j++
	}

	if len(word1) < len(word2) {
		for i := len(word1); i < len(word2); i++ {
			_, _ = sb.WriteString(string(word2[i]))
		}
	} else {
		for i := len(word2); i < len(word1); i++ {
			_, _ = sb.WriteString(string(word1[i]))
		}
	}

	return sb.String()
}

func main() {
	fmt.Println(mergeAlternately("abc", "pqr"))
	fmt.Println(mergeAlternately("ab", "pqrs"))
}
