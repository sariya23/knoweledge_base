package main

import (
	"fmt"
	"runtime"
)

const (
	iterNum = 7
	gourNum = 5
)

func doSmth(in int) {
	for j := 0; j < iterNum; j++ {
		fmt.Printf("i: %d, j: %d\n", in, j)
		runtime.Gosched()
	}
}

func main() {
	for i := 0; i < gourNum; i++ {
		go doSmth(i)
	}
	fmt.Scanln()
}
