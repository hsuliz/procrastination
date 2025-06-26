package leetcode

func maxProfit(prices []int) int {
	var bestProfit int
	minPrice := prices[0]

	for _, price := range prices[1:] {
		profit := price - minPrice

		if minPrice > price {
			minPrice = price
		}

		if profit > bestProfit {
			bestProfit = profit
		}
	}

	return bestProfit
}
