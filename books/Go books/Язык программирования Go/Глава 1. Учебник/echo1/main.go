package main

import (
	"fmt"
	"os"
)

func main() {
	var s string
	for i, v := range os.Args {
		s += fmt.Sprintf("i = %d, %v\n", i, v)
	}
	fmt.Println(s)
}
