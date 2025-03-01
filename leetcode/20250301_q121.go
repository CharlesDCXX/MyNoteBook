package main

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	var minP = prices[0]
	var res int
	for i := 1; i < len(prices); i++ {
		if prices[i] <= minP {
			minP = prices[i]
		} else if prices[i]-minP > res {
			res = prices[i] - minP
		}
	}
	return res
}
func main() {

}
