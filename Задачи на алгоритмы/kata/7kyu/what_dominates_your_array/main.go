package main

import "fmt"

func Dominator(a []int) int {
	buff := map[int]int{}

	for _, v := range a {
		buff[v]++
	}

	for k, v := range buff {
		if v > len(a)/2 {
			return k
		}
	}
	return -1
}

func main() {
	fmt.Println(Dominator([]int{3, 4, 3, 2, 3, 1, 3, 3}))
}
