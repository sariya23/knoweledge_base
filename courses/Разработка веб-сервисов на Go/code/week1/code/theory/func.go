package main

import "fmt"

func main() {
	defer fmt.Println(Smth(), "IN PRINT")
	fmt.Println("WORK")
}

func Smth() string {
	fmt.Println("Smth execute")
	return "s"
}
