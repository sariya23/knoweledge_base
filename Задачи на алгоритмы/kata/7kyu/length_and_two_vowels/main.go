package main

func Alternate(n int, firstValue string, secondValue string) []string {
	result := make([]string, 0, n)
	isAppendFirstValue := true

	for i := 0; i < n; i++ {
		if isAppendFirstValue {
			result = append(result, firstValue)
			isAppendFirstValue = false
		} else {
			result = append(result, secondValue)
			isAppendFirstValue = true
		}
	}

	return result
}
