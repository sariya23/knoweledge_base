package main

func increasingTriplet(nums []int) bool {
	for i := 0; i < len(nums)-2; i++ {
		j, k := i+1, i+2
		if nums[i] < nums[j] && nums[j] < nums[k] && nums[i] < nums[k] {
			return true
		}
	}
	return false
}
