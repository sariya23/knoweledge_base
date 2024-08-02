package main

import "strings"

func Capitalize(st string) []string {
	var evenCapitalize, oddCapitalize string

	for i, v := range st {
		if i%2 == 0 {
			evenCapitalize += strings.ToUpper(string(v))
			oddCapitalize += strings.ToLower(string(v))
		} else {
			evenCapitalize += strings.ToLower(string(v))
			oddCapitalize += strings.ToUpper(string(v))
		}
	}
	return []string{evenCapitalize, oddCapitalize}
}
