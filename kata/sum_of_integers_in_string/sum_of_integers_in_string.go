package main

import (
	"fmt"
	"strconv"
	"strings"
)

func SumOfIntegersInString(strng string) int {
	total := 0
	digits := "0123456789"
	s := ""

	for i := 0; i < len(strng); i++ {
		if strings.Contains(digits, string(strng[i])) {
			s += string(strng[i])
		} else {
			s += "_"
		}
	}

	for _, v := range strings.Split(s, "_") {
		if v != " " {
			n, _ := strconv.Atoi(v)
			total += n
		}
	}

	return total
}

func main() {
	fmt.Println(SumOfIntegersInString("The30quick20brown10f0x1203jumps914ov3r1349the102l4zy dog"))
}
