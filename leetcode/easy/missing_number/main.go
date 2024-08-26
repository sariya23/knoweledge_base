package main

import "fmt"

func missingNumber(nums []int) int {
	a1 := 0
	d := 1
	n := len(nums) + 1
	arithmeticSum := a1*n + d*(((n-1)*n)/2)

	var total int

	for _, v := range nums {
		total += v
	}

	return arithmeticSum - total

}

func main() {
	fmt.Println(missingNumber([]int{1}))
}
