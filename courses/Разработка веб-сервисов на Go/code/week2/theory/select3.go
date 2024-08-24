package main

import "fmt"

func main() {
	done := make(chan struct{})
	ch := make(chan int)

	go func() {
		val := 0
		for {
			select {
			case <-done:
				return
			case ch <- val:
				val++
				fmt.Println(val)
			}
		}
	}()

	for c := range ch {
		fmt.Println("read")
		if c > 3 {
			fmt.Println("send cancel")
			close(done)
			break
		}
	}
}
