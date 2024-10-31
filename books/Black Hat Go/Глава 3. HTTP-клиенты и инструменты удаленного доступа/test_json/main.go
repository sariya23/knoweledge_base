package main

import (
	"fmt"
	"unsafe"
)

type s1 struct {
	d string
	b bool
	a string
	s bool
}

type s2 struct {
	d []string
	// a string
	// b bool
	// s bool
}

func main() {
	x := s1{}
	y := s2{}
	fmt.Println("s1", unsafe.Sizeof(x)) // 48
	fmt.Println("s2", unsafe.Sizeof(y)) // 40
}
