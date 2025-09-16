package main

func maxProfit(prices []int) int {
	maxProfit := 0
	cheapest := 10_000

	for _, v := range prices {
		if v < cheapest {
			cheapest = v
		}

		if currProfit := v - cheapest; currProfit > maxProfit {
			maxProfit = currProfit
		}
	}

	return maxProfit
}
