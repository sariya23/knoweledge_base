package main

import (
	"fmt"
	"math"
)

func TwoOldestAges(ages []int) [2]int {
	max := math.MinInt + 1
	preMax := math.MinInt + 1

	for _, v := range ages {
		if v > max {
			preMax = max
			max = v
		} else if v > preMax && v <= max {
			preMax = v
		}
	}
	return [...]int{preMax, max}
}

func main() {
	fmt.Println(TwoOldestAges([]int{6, 5, 83, 5, 3, 18}))                     // 18 83
	fmt.Println(TwoOldestAges([]int{39, 53, 83, 51, 59, 61, 95, 23, 99, 49})) // 95 99
	fmt.Println(TwoOldestAges([]int{93, 35, 53, 67, 17, 23, 89, 75, 15, 53})) // 89 93
	fmt.Println(TwoOldestAges([]int{19, 5, 43, 13, 75, 89, 43, 89, 25, 49}))  // 75 89
}
