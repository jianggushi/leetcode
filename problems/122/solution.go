package solution

// 时间复杂度 O(n) 空间复杂度 O(1)
// 这次可以多次买卖，策略肯定还是低买高卖，在低位买入，高位卖出
func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	ans := 0
	in, out := prices[0], prices[0] // 初始买入点和卖出点
	for i := 1; i < n; i++ {
		if prices[i] > out { // 上升过程中，更新卖出点
			out = prices[i]
		} else {
			ans += out - in // 到了高位，卖出，更新买入点和卖出点
			in = prices[i]
			out = prices[i]
		}
	}
	return ans + (out - in) // !尾巴上是上升过程，把最后的卖出
}
