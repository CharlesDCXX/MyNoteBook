package main

func maxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}
	var minV = prices[0]
	var re = 0
	for i := 1; i < len(prices); i++ {
		if minV < prices[i] {
			re += prices[i] - minV
		}
		minV = prices[i]
	}
	return re
}
