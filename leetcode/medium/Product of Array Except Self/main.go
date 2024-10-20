package main

import "fmt"

func productExceptSelf(nums []int) []int {
	res := make([]int, 0, len(nums))
	left, right := make([]int, len(nums)), make([]int, len(nums))
	left[0] = 1
	for i := 1; i < len(nums); i++ {
		left[i] = left[i-1] * nums[i-1]
	}

	right[len(nums)-1] = 1
	for i := len(nums) - 2; i > -1; i-- {
		right[i] = right[i+1] * nums[i+1]
	}

	for i := range left {
		res = append(res, left[i]*right[i])
	}

	return res
}

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}
