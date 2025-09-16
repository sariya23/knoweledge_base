package main

func findInvertedInteger(arr []int, value int) bool {
	for _, v := range arr {
		if v == -value {
			return true
		}
	}
	return false
}
func Solve(arr []int) int {
	var result int
	for _, v := range arr {
		if !findInvertedInteger(arr, v) {
			result = v
		}
	}
	return result
}
