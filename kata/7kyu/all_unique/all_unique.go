package main

func HasUniqueChar(str string) bool {
	set := map[string]struct{}{}

	for _, v := range str {
		set[string(v)] = struct{}{}
	}
	return len(str) == len(set)
}
