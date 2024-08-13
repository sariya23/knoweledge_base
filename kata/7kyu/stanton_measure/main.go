package main

func Count(arr []int, value int) int {
	var counter int

	for _, v := range arr {
		if v == value {
			counter++
		}
	}

	return counter
}

func StantonMeasure(arr []int) int {
	amountOfOne := Count(arr, 1)
	return Count(arr, amountOfOne)
}
