package main

import (
	"fmt"
	"strings"
)

func makeMatrixFromString(s string) [][]string {
	splittedStrings := strings.Split(s, "\n")
	matrix := make([][]string, len(string(splittedStrings[0])))
	for i := 0; i < len(splittedStrings); i++ {
		row := make([]string, len(string(splittedStrings[i])))
		for j := 0; j < len(string(splittedStrings[i])); j++ {
			row[j] = string(splittedStrings[i][j])
		}
		matrix[i] = row
	}
	return matrix
}

func VertMirror(s string) string {
	maxtrix := makeMatrixFromString(s)
	result := make([]string, len(maxtrix))

	for i := 0; i < len(maxtrix); i++ {
		elem := ""
		for j := len(maxtrix[i]) - 1; j > -1; j-- {
			elem += string(maxtrix[i][j])
		}

		result[i] = elem
	}

	return strings.Join(result, "\n")
}
func HorMirror(s string) string {
	maxtrix := makeMatrixFromString(s)
	result := make([]string, len(maxtrix))
	for i := len(maxtrix) - 1; i > -1; i-- {
		result[len(maxtrix)-i-1] = strings.Join(maxtrix[i], "")
	}
	return strings.Join(result, "\n")
}

type FParam func(string) string

func Oper(f FParam, x string) string {
	return f(x)
}

func main() {
	s := "abcd\nefgh\nijkl\nmnop"
	fmt.Println(HorMirror(s))
	// fmt.Println(Oper(HorMirror, s))
}
