package main

import (
	"fmt"
	"strings"
)

func RemoveDuplicateWords(str string) string {
	words := strings.Split(str, " ")
	uniqueWords := make([]string, 0)
	counter := map[string]int{}

	for _, v := range words {
		if string(v) == "" {
			continue
		}
		_, ok := counter[string(v)]

		if !ok {
			uniqueWords = append(uniqueWords, string(v))
			counter[string(v)]++
		}
	}
	return strings.Join(uniqueWords, " ")
}

func main() {
	fmt.Println(RemoveDuplicateWords("my cat is    my cat fat"))
}
