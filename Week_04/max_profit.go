package week04

func maxProfit(prices []int) int {
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-prices[i-1] > 0 { // 比较今天和昨天的价格,今天比昨天的价格高就卖出
			profit += prices[i] - prices[i-1] // 存储利润
		}
	}
	return profit
}
