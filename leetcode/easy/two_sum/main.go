package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i, v := range nums {
		m[target-v] = i
	}
	fmt.Println(m)
	res := make([]int, 2)
	for i, v := range nums {
		if j, ok := m[v]; ok && j != i {
			res[0] = i
			res[1] = j
			break
		}
	}
	return res
}

func main() {
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
}
