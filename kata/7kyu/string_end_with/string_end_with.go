package main

func ReverseString(s string) string {
	result := ""

	for i := len(s) - 1; i >= 0; i-- {
		result += string(s[i])
	}
	return result
}

func solution(str, ending string) bool {
	if len(str) < len(ending) {
		return false
	}
	s := ReverseString(str)
	sEnd := ReverseString(ending)

	for i := 0; i < len(sEnd); i++ {
		if s[i] != sEnd[i] {
			return false
		}
	}

	return true
}
