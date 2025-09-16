package main

import "fmt"

func GrowingPlant(upSpeed, downSpeed, desiredHeight int) int {
	var currentHeight, remainingDays int

	for {
		remainingDays++
		currentHeight += upSpeed
		if currentHeight >= desiredHeight {
			break
		}
		currentHeight -= downSpeed
	}
	return remainingDays
}

func main() {
	fmt.Println(GrowingPlant(100, 10, 910))
	fmt.Println(GrowingPlant(10, 9, 4))
}
