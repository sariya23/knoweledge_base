package main

import "fmt"

func RotateString(s string, relativeRotateIndex int) string {
	withAfterSlice := s[relativeRotateIndex:]
	beforeSlice := s[:relativeRotateIndex]
	return withAfterSlice + beforeSlice
}

func ContainAllRots(strng string, arr []string) bool {
	if strng == "" {
		return true
	}
	amountOfRotate := len(strng)

	for i := 0; i < amountOfRotate; i++ {
		rotated := RotateString(strng, i)
		found := false
		for _, v := range arr {
			if v == rotated {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(ContainAllRots("bsjq", []string{"bsjq", "qbsj", "sjqb", "twZNsslC", "jqbs"}))
	fmt.Println(ContainAllRots("XjYABhR", []string{"TzYxlgfnhf", "yqVAuoLjMLy", "BhRXjYA", "YABhRXj", "hRXjYAB", "jYABhRX", "XjYABhR", "ABhRXjY"}))
	fmt.Println(ContainAllRots("", []string{}))
	fmt.Println(ContainAllRots("Ajylvpy", []string{"Ajylvpy", "ylvpyAj", "jylvpyA", "lvpyAjy", "pyAjylv", "vpyAjyl", "ipywee"}))
}
