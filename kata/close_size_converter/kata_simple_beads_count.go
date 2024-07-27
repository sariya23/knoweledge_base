package main

import (
	"fmt"
	"strings"
)

func ValidSize(size string) bool {
	if len(size) == 0 {
		return false
	}

	if string(size[0]) != "x" {
		fmt.Printf("wrong 1 symbol %s\n", string(size[0]))
		return false
	}

	if string(size[len(size)-1]) != "l" && string(size[len(size)-1]) != "s" {
		fmt.Printf("wrong last symbol %s\n", string(size[len(size)-1]))
		return false
	}

	for i := 0; i < len(size)-1; i++ {
		if string(size[i]) != "x" {
			return false
		}
	}

	return true
}

func SizeToNumber(size string) (int, bool) {
	if size == "s" {
		return 36, true
	} else if size == "m" {
		return 38, true
	} else if size == "l" {
		return 40, true
	} else {
		if ValidSize(size) {
			if strings.Contains(size, "s") {
				return 36 - (strings.Count(size, "x") * 2), true
			} else if strings.Contains(size, "l") {
				return 40 + (strings.Count(size, "x") * 2), true
			} else {
				return 0, false
			}
		}
	}
	return 0, false
}

func main() {
	fmt.Println(SizeToNumber("ssss"))
}
