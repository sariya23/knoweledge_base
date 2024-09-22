package main

import (
	"fmt"
	"strings"
)

func gcdOfStrings(str1 string, str2 string) string {
	type res struct {
		Val string
		Len int
	}

	gcd := res{}

	i, j := 0, 0
	var currGcd string

	for i < len(str1) && j < len(str2) {
		if str1[i] == str2[j] {
			currGcd += string(str1[i])
		}
		if canBeGCDOfString(str1, currGcd) && canBeGCDOfString(str2, currGcd) {
			if len(currGcd) > gcd.Len {
				gcd.Len = len(currGcd)
				gcd.Val = currGcd
			}
		}
		i++
		j++

	}
	return gcd.Val
}

func canBeGCDOfString(src string, gcd string) bool {
	if len(gcd) == 0 {
		return false
	}
	if len(src)%len(gcd) != 0 {
		return false
	}
	total := len(src) / len(gcd)
	sb := strings.Builder{}
	sb.Grow(total)
	for i := 0; i < total; i++ {
		_, _ = sb.WriteString(gcd)
	}
	return sb.String() == src
}

func main() {
	fmt.Println(gcdOfStrings("ABCABC", "ABC"))
	fmt.Println(gcdOfStrings("ABABAB", "ABAB"))
	fmt.Println(gcdOfStrings("LEET", "CODE"))
	fmt.Println(gcdOfStrings("ABABABAB", "ABAB"))
}
