package main

import "fmt"

func containsDuplicate(nums []int) bool {
	set := make(map[int]struct{})

	for _, v := range nums {
		set[v] = struct{}{}
	}

	return len(nums) != len(set)
}

func main() {
	s := make(map[int]struct{})
	s[2] = struct{}{}
	s[2] = struct{}{}
	fmt.Println(s)
}
