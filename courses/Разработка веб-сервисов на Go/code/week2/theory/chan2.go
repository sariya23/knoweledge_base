package main

import "fmt"

func main() {
	in := make(chan int)

	go func(in chan<- int) {
		for i := 0; i < 5; i++ {
			in <- i
		}
		close(in)
	}(in)
	for i := range in {
		fmt.Println(i)
	}
}
