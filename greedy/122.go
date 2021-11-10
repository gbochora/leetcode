package greedy

func maxProfit(prices []int) int {
	var i, totalProfit, localProfit int
	for i < len(prices) {
		i, localProfit = localMaxProfit(prices, i)
		totalProfit += localProfit
	}
	return totalProfit
}

func localMaxProfit(prices []int, from int) (pos, profit int) {
	initPrice := prices[from]
	for from+1 < len(prices) && prices[from] <= prices[from+1] {
		from++
	}
	return from + 1, prices[from] - initPrice
}
