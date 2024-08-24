package main

import (
	"fmt"
	"sync"
)

func deferTest() {
	a, b := 1, 0
	fmt.Println(a / b)
	defer fmt.Println("end")
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(i)
		}()
	}
	wg.Wait()
}
