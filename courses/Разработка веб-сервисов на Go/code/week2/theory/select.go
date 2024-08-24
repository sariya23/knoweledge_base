package main

import "fmt"

func main() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int)

	ch1 <- 2

	select {
	case val := <-ch1:
		fmt.Println("ch1", val)
	case ch2 <- 1:
		fmt.Println("put val in ch2")
	default:
		fmt.Println("default")
	}
}
