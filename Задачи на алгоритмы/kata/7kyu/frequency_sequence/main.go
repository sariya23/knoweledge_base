package main

import (
	"strconv"
	"strings"
)

func FreqSeq(str string, sep string) string {
	result := make([]string, 0, len(str)+len(str)-1)

	for _, v := range str {
		result = append(result, strconv.Itoa(strings.Count(str, string(v))))
	}

	return strings.Join(result, sep)
}
