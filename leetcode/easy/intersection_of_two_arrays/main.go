package main

import (
	"fmt"
	"slices"
)

// nlog(n) + mlog(m) + m*log(m)
func intersection(nums1 []int, nums2 []int) []int {
	slices.Sort(nums1)
	slices.Sort(nums2)

	var bigNums, smallNums []int

	if len(nums1) > len(nums2) {
		bigNums = nums1
		smallNums = nums2
	} else {
		bigNums = nums2
		smallNums = nums1
	}

	buf := make(map[int]struct{})

	var result []int

	for i := 0; i < len(smallNums); i++ {
		inNum1 := binarySeacrh(bigNums, smallNums[i]) != -1

		if _, ok := buf[smallNums[i]]; !ok && inNum1 {
			buf[smallNums[i]] = struct{}{}
		}
	}

	for k := range buf {
		result = append(result, k)
	}

	return result
}

func binarySeacrh(arr []int, target int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := arr[mid]

		if target == guess {
			return mid
		} else if guess > target {
			high = mid - 1
		} else {
			low = high - 1
		}
	}
	return -1
}

func main() {
	fmt.Println(intersection([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Println()
	fmt.Println(intersection([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
}
