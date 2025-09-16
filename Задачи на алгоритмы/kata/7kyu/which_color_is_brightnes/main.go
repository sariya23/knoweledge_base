package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func parseHexColorToInt(color string) (int, int, int) {
	colorNoSharp := strings.Replace(color, "#", "", 1)
	n1, _ := strconv.ParseInt(colorNoSharp[:2], 16, 64)
	n2, _ := strconv.ParseInt(colorNoSharp[2:4], 16, 64)
	n3, _ := strconv.ParseInt(colorNoSharp[4:], 16, 64)

	return int(n1), int(n2), int(n3)
}

type Color struct {
	Hex string
	V   int
}

func Brightest(colors []string) string {

	maxColor := Color{
		Hex: "",
		V:   0,
	}

	for _, v := range colors {
		r, g, b := parseHexColorToInt(string(v))
		fmt.Printf("r %d g %d b %d \t max color: ", r, g, b)
		V := int(math.Max(math.Max(float64(r), float64(g)), float64(b)))
		fmt.Println(V)

		if V > maxColor.V {
			maxColor.Hex = string(v)
			maxColor.V = V
		}
	}
	return maxColor.Hex
}

func main() {
	fmt.Println(Brightest([]string{"#115A53", "#AC1672", "#0D8C3D", "#2BA73A", "#58F38F", "#F27CFD", "#8AEBF9"}))
}
