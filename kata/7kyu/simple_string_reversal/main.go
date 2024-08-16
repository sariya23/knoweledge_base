package main

import (
	"fmt"
	"strings"
)

func GetSpaceIndexes(s string) []int {
	res := make([]int, 0, strings.Count(s, " "))
	for i := 0; i < len(s); i++ {
		if string(s[i]) == " " {
			res = append(res, i)
		}
	}

	return res
}

func reverseString(s string) string {
	var result string

	for i := len(s) - 1; i > -1; i-- {
		result += string(s[i])
	}
	return result
}

func insert(s string, v string, i int) string {
	return s[:i] + v + s[i:]
}

func Solve(s string) string {
	spaceIndexes := GetSpaceIndexes(s)
	reversedString := reverseString(s)
	reversedStringNoSpaces := strings.ReplaceAll(reversedString, " ", "")

	for _, v := range spaceIndexes {
		reversedStringNoSpaces = insert(reversedStringNoSpaces, " ", v)
	}

	return reversedStringNoSpaces
}

func main() {
	fmt.Println(Solve("our code"))
	fmt.Println(Solve("your code rocks"))
	fmt.Println(Solve("codewars"))
}
