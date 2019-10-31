package solution

// Say you have an array for which the ith element is the price of a given stock on day i.
// If you were only permitted to complete at most one transaction (i.e., buy one and sell one share of the stock), design an algorithm to find the maximum profit.
// Note that you cannot sell a stock before you buy one.

// Example 1:
// Input: [7,1,5,3,6,4]
// Output: 5
// Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
//              Not 7-1 = 6, as selling price needs to be larger than buying price.

// Example 2:
// Input: [7,6,4,3,1]
// Output: 0
// Explanation: In this case, no transaction is done, i.e. max profit = 0.

// 暴力搜索
func maxProfit1(prices []int) int {
	n := len(prices)
	result := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			tmp := prices[j] - prices[i]
			if tmp > result {
				result = tmp
			}
		}
	}
	return result
}

func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	in, out := prices[0], prices[0]
	result := 0

	for i := 1; i < n; i++ {
		if prices[i] > out {
			out = prices[i]
			tmp := out - in
			if tmp > result {
				result = tmp
			}
		} else if prices[i] < in {
			in = prices[i]
			out = in
		}
	}
	return result
}
