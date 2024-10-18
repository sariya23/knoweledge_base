package main

func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxCandies := max(candies)
	res := make([]bool, 0, len(candies))

	for _, v := range candies {
		if v+extraCandies >= maxCandies {
			res = append(res, true)
		} else {
			res = append(res, false)
		}
	}
	return res
}

func max(arr []int) int {
	m := -1
	for _, v := range arr {
		if v > m {
			m = v
		}
	}
	return m
}
