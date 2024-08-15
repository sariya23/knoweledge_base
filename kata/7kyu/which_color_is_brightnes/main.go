package main

import (
	"fmt"
	"strconv"
	"strings"
)

func parseHexColorToInt(color string) []int {
	colorNoSharp := strings.Replace(color, "#", "", 1)
	n1, _ := strconv.ParseInt(colorNoSharp[:2], 16, 8)
	n2, _ := strconv.ParseInt(colorNoSharp[2:4], 16, 8)
	n3, _ := strconv.ParseInt(colorNoSharp[4:], 16, 8)

	return []int{int(n1), int(n2), int(n3)}
}

type Color struct {
	hex        string
	decimalRep int
}

func Brightest(colors []string) string {
	maxColor := Color{
		hex:        "",
		decimalRep: -1,
	}

	for _, v := range colors {
		colorHex := strings.Replace(string(v), "#", "", 1)
		color, err := strconv.ParseInt(colorHex, 16, 64)
		if err != nil {
			panic(err)
		}
		if color > int64(maxColor.decimalRep) {
			maxColor.decimalRep = int(color)
			maxColor.hex = string(v)
		}
	}

	return maxColor.hex
}

func main() {
	for _, v := range []string{"#00FF00", "#FFFF00", "#01130F"} {
		color := strings.Replace(string(v), "#", "", 1)
		n1, _ := strconv.ParseInt(color[:2], 16, 8)
		n2, _ := strconv.ParseInt(color[2:4], 16, 8)
		n3, _ := strconv.ParseInt(color[4:], 16, 8)
		fmt.Println(n1, n2, n3)
	}
}
