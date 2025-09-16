package main

import (
	"fmt"
	"sort"
)

type Runes []rune

func (r Runes) Len() int {
	return len(r)
}

func (r Runes) Less(i, j int) bool {
	return r[i] < r[j]
}

func (r Runes) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func Solve(s string) bool {
	runes := Runes([]rune(s))
	sort.Sort(runes)

	for i := 0; i < len(runes)-1; i++ {
		if runes[i+1]-runes[i] != 1 {
			return false
		}
	}

	return true
}

func main() {
	r := Runes([]rune("cba"))
	sort.Sort(r)
	fmt.Println(r)
	// fmt.Println(Solve("abc"))
	// fmt.Println(Solve("abd"))
	// fmt.Println(Solve("dabc"))
	// fmt.Println(Solve("abc"))
	// fmt.Println(Solve("abc"))
}
