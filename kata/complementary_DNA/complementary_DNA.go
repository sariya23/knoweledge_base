package main

import "fmt"

func DNAStrand(dna string) string {
	dna_map := map[string]string{
		"A": "T",
		"C": "G",
		"T": "A",
		"G": "C",
	}
	sideDNA := ""

	for _, v := range dna {
		fmt.Println(string(v))
		sideDNA += dna_map[string(v)]
	}

	return sideDNA
}

func main() {
	fmt.Println(DNAStrand("ATTGC"))
}
