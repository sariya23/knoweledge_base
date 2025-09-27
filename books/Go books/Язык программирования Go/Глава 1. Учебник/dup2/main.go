package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := map[string]map[string]int{}
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, "")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			name := f.Name()
			defer f.Close()
			countLines(f, counts, name)
		}
	}
	for k, v := range counts {
		fmt.Printf("strings %v in files:\n", k)
		for f, n := range v {
			if n > 1 {
				fmt.Printf("\tfile %v, %d times\n", f, n)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]map[string]int, name string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		if _, ok := counts[line]; !ok {
			counts[line] = make(map[string]int)
		}
		counts[line][name]++
	}

}
