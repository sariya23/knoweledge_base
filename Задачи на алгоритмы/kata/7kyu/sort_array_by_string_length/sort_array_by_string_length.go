package main

import "fmt"

func lessPivot(arr []string, pivot string) []string {
	result := make([]string, 0)

	for _, v := range arr {
		if len(string(v)) < len(pivot) {
			result = append(result, string(v))
		}
	}
	return result
}

func greaterPivot(arr []string, pivot string) []string {
	result := make([]string, 0)

	for _, v := range arr {
		if len(string(v)) > len(pivot) {
			result = append(result, string(v))
		}
	}
	return result
}

func QuickSort(arr []string) []string {

	if len(arr) <= 1 {
		return arr
	}

	pivot := string(arr[0])
	lessPivot := lessPivot(arr, pivot)
	greaterPivot := greaterPivot(arr, pivot)

	return append(append(QuickSort(lessPivot), []string{pivot}...), QuickSort(greaterPivot)...)
}

func SortByLength(arr []string) []string {
	return QuickSort(arr)
}

func main() {
	fmt.Println(QuickSort([]string{"beg", "life", "i", "to"}))
}
