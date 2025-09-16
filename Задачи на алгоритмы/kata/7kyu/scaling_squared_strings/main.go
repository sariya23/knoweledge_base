package main

import (
	"fmt"
	"strings"
)

func verticalScale(s string, replicate int) string {
	var result string

	for _, v := range s {
		result += strings.Repeat(string(v), replicate)
	}

	return result
}

func horizontalScale(s string, replicate int) []string {
	result := make([]string, 0, replicate)

	for i := 0; i < replicate; i++ {
		result = append(result, s)
	}

	return result
}

func Scale(s string, k, n int) string {
	if s == "" {
		return ""
	}
	result := make([]string, 0, len(s)*n)
	rows := strings.Split(s, "\n")
	for _, r := range rows {
		tempRow := horizontalScale(verticalScale(string(r), k), n)
		result = append(result, tempRow...)
	}
	return strings.Join(result, "\n")
}

func main() {
	fmt.Println(Scale("", 2, 3))
}
