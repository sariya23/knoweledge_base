package main

import (
	"fmt"
	"strings"
)

func sliceContain(arr []int, target int) bool {
	for _, v := range arr {
		if v == target {
			return true
		}
	}
	return false
}

func GetSpaceIndexes(s string) []int {
	res := make([]int, 0, strings.Count(s, " "))
	for i := 0; i < len(s); i++ {
		if string(s[i]) == " " {
			res = append(res, i)
		}
	}

	return res
}

func reverseStringAndExcludeSpaces(s string) string {
	var result string

	for i := len(s) - 1; i > -1; i-- {
		if v := string(s[i]); v != " " {
			result += v
		}
	}
	return result
}

func Solve(s string) string {
	var result string
	spaceIndexes := GetSpaceIndexes(s)
	reversedStringNoSpaces := reverseStringAndExcludeSpaces(s)
	fmt.Println(reversedStringNoSpaces)

	for i := 0; i < len(reversedStringNoSpaces); i++ {
		if sliceContain(spaceIndexes, i) {
			result += fmt.Sprintf(" %s", string(reversedStringNoSpaces[i]))
		} else {
			result += string(reversedStringNoSpaces[i])
		}
	}

	return result
}

func main() {
	fmt.Println(Solve("our code"))
	fmt.Println(Solve("your code rocks"))
	fmt.Println(Solve("srawedoc"))
}
