package main

import "fmt"

func lettersAtEvenIndexes(s string) string {
	var result string

	for i, v := range s {
		if i%2 == 0 {
			result += string(v)
		}
	}
	return result
}

func lettersAtOddIndexes(s string) string {
	var result string

	for i, v := range s {
		if i%2 != 0 {
			result += string(v)
		}
	}
	return result
}

func SortMyString(s string) string {
	return fmt.Sprintf("%s %s", lettersAtEvenIndexes(s), lettersAtOddIndexes(s))
}
