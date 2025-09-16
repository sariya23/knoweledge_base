package main

import (
	"fmt"
	"math"
)

func calcAvarageHourSpeed(delta_distance float64, s int) float64 {
	return float64(3600*delta_distance) / float64(s)
}

func findMax(x []float64) float64 {
	m := x[0]

	for _, v := range x {
		if v > m {
			m = v
		}
	}

	return m
}

func Gps(s int, x []float64) int {
	if len(x) <= 1 {
		return 0
	}

	averageSpeed := make([]float64, 0, len(x)-1)

	for i := 0; i < len(x)-1; i++ {
		averageSpeed = append(averageSpeed, calcAvarageHourSpeed(x[i+1]-x[i], s))
	}

	maxSpeed := findMax(averageSpeed)

	return int(math.Floor(maxSpeed))
}

func main() {
	fmt.Println(Gps(15, []float64{0.0, 0.19, 0.5, 0.75, 1.0, 1.25, 1.5, 1.75, 2.0, 2.25}))
}
