package main

import "fmt"

func PickDifferentColor(color1, color2 string) string {
	m := map[string]rune{
		"R": 82,
		"G": 71,
		"B": 66,
	}
	letterColorCheckSum := m["R"] + m["G"] + m["B"]
	colorSum := m[color1] + m[color2]
	return string(letterColorCheckSum - colorSum)
}

func Triangle(row string) rune {
	if len(row) == 1 {
		return rune(row[0])
	}
	rowAmount := len(row) - 1

	for i := 0; i < rowAmount; i++ {
		var newRow string
		prevRow := row
		for j := 0; j < len(prevRow)-1; j++ {
			if prevRow[j] == prevRow[j+1] {
				newRow += string(prevRow[j])
			} else {
				newRow += PickDifferentColor(string(prevRow[j]), string(prevRow[j+1]))
			}
		}
		row = newRow
	}
	return rune(row[0])
}

func main() {
	fmt.Println(Triangle("GB"))
	fmt.Println(Triangle("RRR"))
	fmt.Println(Triangle("RGBG"))
	fmt.Println(Triangle("RBRGBRBGGRRRBGBBBGG"))
	fmt.Println(Triangle("B"))
}
