package main

import (
	"fmt"
	"math"
	"strings"
)

// Solve возвращает самую длинную неприрывную цепочку
// из гласных букв английского алфавита.
// Принимает один параметр - строку-слово. Содержит только
// буквы латинского алфавита в нижнем регистре.
func Solve(s string) int {
	const vowels = "aeiou"
	var bufferVowels int
	var maxChainLen int

	for _, v := range s {
		if strings.Contains(vowels, string(v)) {
			bufferVowels++
		} else {
			maxChainLen = int(math.Max(float64(maxChainLen), float64(bufferVowels)))
			bufferVowels = 0
		}
	}

	return maxChainLen
}

func main() {
	fmt.Println(Solve("codewarriors"))
}
