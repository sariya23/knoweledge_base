package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	words := strings.Split(s, " ")
	words = reverseSlice(words)
	words = deleteSpacesFromSlice(words)
	return strings.Join(words, " ")
}

func reverseSlice(arr []string) []string {
	res := make([]string, 0, len(arr))

	for i := len(arr) - 1; i > -1; i-- {
		res = append(res, arr[i])
	}
	return res
}

func deleteSpacesFromSlice(arr []string) []string {
	res := make([]string, 0, len(arr))
	for _, v := range arr {
		if v != "" {
			res = append(res, v)
		}
	}
	return res
}

func main() {
	fmt.Println(reverseWords("a good   example"))
	fmt.Println(reverseSlice([]string{"a", "b", "cv"}))
	fmt.Println(deleteSpacesFromSlice([]string{"a", "", "v", ""}))
	fmt.Println(strings.Fields("q        b"))
}
