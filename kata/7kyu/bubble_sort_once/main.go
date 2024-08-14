package main

func BubblesortOnce(numbers []int) []int {
	result := make([]int, len(numbers))
	copy(result, numbers)
	for i := 0; i < len(result)-1; i++ {
		if result[i] > result[i+1] {
			result[i], result[i+1] = result[i+1], result[i]
		}
	}
	return result
}
