package main

import "math"

func findStringWithMaxLen(s []string) string {
	max := string(s[0])

	for _, v := range s[1:] {
		if len(string(v)) > len(max) {
			max = string(v)
		}
	}

	return max
}

func findStringWithMinLen(s []string) string {
	min := string(s[0])

	for _, v := range s[1:] {
		if len(string(v)) < len(min) {
			min = string(v)
		}
	}

	return min
}

func getABSDifferenceBetweenStringLens(s1, s2 string) int {
	return int(math.Abs(float64(len(s1) - len(s2))))
}

func MxDifLg(a1 []string, a2 []string) int {
	if len(a1) == 0 || len(a2) == 0 {
		return -1
	}

	maxLenStringA1 := findStringWithMaxLen(a1)
	minLenStringA1 := findStringWithMinLen(a1)

	maxLenStringA2 := findStringWithMaxLen(a2)
	minLenStringA2 := findStringWithMinLen(a2)

	diffA1A2 := getABSDifferenceBetweenStringLens(maxLenStringA1, minLenStringA2)
	diffA2A1 := getABSDifferenceBetweenStringLens(maxLenStringA2, minLenStringA1)

	if diffA1A2 > diffA2A1 {
		return diffA1A2
	} else {
		return diffA2A1
	}
}
