package main

import (
	"fmt"
	"unsafe"
)

func main() {
	n := 10
	fmt.Printf("%T\n", n)
	fmt.Println(unsafe.Sizeof(n))

	d := new(int)
	*d = 10
	fmt.Println(*d)

	var l []int
	sl := make([]int, 0)
	fmt.Println(l, sl)
	fmt.Println(l == nil)
	fmt.Println(sl == nil)
	s := make([]int, 0, 2)
	s = append(s, []int{1, 2}...)
	fmt.Println(s)

	s1 := s[:]
	s1 = append(s1, 3)
	fmt.Println(s)
	s1[0] = 10
	fmt.Println(s1)

	for i := range []int{1, 2, 3} {
		fmt.Println(i)
	}

	for k := range map[string]int{"a": 1, "b": 2} {
		fmt.Println(k)
	}
}
