package main

import "fmt"

func contain(arr []int, value int) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}

func Solve(arr []int) []int {
	result := make([]int, 0, len(arr))

	for i, v := range arr {
		if !contain(arr[i+1:], v) {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	fmt.Println(Solve([]int{3, 4, 4, 3, 6, 3}))
}
