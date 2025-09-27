package main

import "fmt"

func main() {
	for i := 1; i < 25; i++ {
		fmt.Println(1 << i)
	}
	a := 1.99
	fmt.Println(int(a))
}
