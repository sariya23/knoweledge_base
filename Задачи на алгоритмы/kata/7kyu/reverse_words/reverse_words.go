package main

import (
	"fmt"
	"strings"
)

func ReverseString(str string) string {
	reversed := ""

	for i := len(str) - 1; i > -1; i-- {
		reversed += string(str[i])
	}

	return reversed
}

func ReverseWords(str string) string {
	words := strings.Split(str, " ")
	reversedWords := make([]string, 0)

	for _, s := range words {
		reversedWords = append(reversedWords, ReverseString(string(s)))
	}

	return strings.Join(reversedWords, " ")
}

func main() {
	fmt.Println(ReverseWords("The quick brown fox jumps over the lazy dog."))
}
