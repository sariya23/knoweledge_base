package main

import (
	"fmt"
	"slices"
)

func targetIndices(nums []int, target int) []int {
	slices.Sort(nums)
	l := 0
	r := len(nums)
	var indicies []int
	for l < r {
		if nums[l] == target {
			indicies = append(indicies, nums[l])
		}
	}
	return indicies
}

func f(nums []int, target int) []int {
	targetCount := 0
	smallerCoount := 0

	for _, v := range nums {
		if v == target {
			targetCount++
		} else if v < target {
			smallerCoount++
		}
	}

	fmt.Println("target:", targetCount)
	fmt.Println("smaller:", smallerCoount)

	var res []int

	for i := 0; i < targetCount; i++ {
		res = append(res, smallerCoount)
		smallerCoount++
	}

	return res
}

func main() {
	fmt.Println(f([]int{10, 100, -1, -10, 1, 2, 5, 2, 3}, 2)) // 1 2 2 3 5
}
