package main

import (
	"fmt"
	"strings"
)

func bandNameGenerator(word string) string {
	var bandName string

	if word[0] == word[len(word)-1] {
		suffix := word[1:]
		bandName = strings.Title(word) + suffix
	} else {
		bandName = fmt.Sprintf("The %s", strings.Title(word))
	}
	return bandName
}

func main() {
	fmt.Println(bandNameGenerator("dolphin"))
}
