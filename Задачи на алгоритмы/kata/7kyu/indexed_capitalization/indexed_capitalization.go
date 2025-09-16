package main

import "strings"

func Capitalize(st string, arr []int) string {
	stringLen := len(st)

	for _, v := range arr {
		if v >= stringLen {
			st = strings.Replace(st, string(st[stringLen-1]), strings.ToUpper(string(st[stringLen-1])), 1)
		} else {
			st = strings.Replace(st, string(st[v]), strings.ToUpper(string(st[v])), 1)
		}
	}
	return st
}

func main() {
	println(Capitalize("abcdef", []int{1, 2, 5}))
}
