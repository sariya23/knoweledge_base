package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := os.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, l := range strings.Split(string(data), "\n") {
			counts[l]++
		}
	}

	for l, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, l)
		}
	}
}
