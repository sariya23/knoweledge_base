package main

import (
	"fmt"
)

func main() {
	s := []int{1, 1, 2, 2, 3, 5, 6}
	fmt.Println(removeDuplicates(s))
	fmt.Println(s)
}

func removeDuplicates(nums []int) int {
	m := make(map[int]int, len(nums))
	for _, v := range nums {
		m[v]++
	}
	fmt.Println(m)
	var k int

	for _, v := range m {
		if v == 1 {
			k++
		} else {
			k += v - (v - 1)
		}
	}
	return k
}

func binSearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := arr[mid]

		if guess == target {
			return mid
		} else if guess < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
