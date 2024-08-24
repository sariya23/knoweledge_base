package main

import "fmt"

func main() {
	ch1 := make(chan int, 1)

	go func(ch chan int) {
		val := <-ch
		fmt.Println("GO: get from chan1", val)
		fmt.Println("GO: After read")
	}(ch1)
	ch1 <- 42
	ch1 <- 1
	fmt.Println("MAIN: after write")
	fmt.Println(<-ch1)
	fmt.Scanln()
}
