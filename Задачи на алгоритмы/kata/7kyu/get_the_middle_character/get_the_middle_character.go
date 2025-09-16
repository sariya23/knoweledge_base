package main

import "fmt"

func GetMiddle(s string) string {
	result := ""

	if len(s)%2 == 0 {
		result = string(s[len(s)/2-1]) + string(s[len(s)/2])
	} else {
		result = string(s[len(s)/2])
	}
	return result
}

func main() {
	fmt.Println(GetMiddle("test"))
}
