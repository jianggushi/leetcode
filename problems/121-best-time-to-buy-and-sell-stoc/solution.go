// Source : https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/
// Author : jianggushi
// Date   : 2022-11-15

package solution

// 优化二
// 在每一个位置都计算卖出的收益，但有时候的计算是无意义的
// 比如在下降阶段，买入点一直在更新，收益在无意义的更新检查
// 为了避免检查一些无意义的卖出点，要对之前的卖出点进行记录，只有当前价格大于之前的卖出点再更新卖出点和收益
// ! 在下降阶段，买入点一直在更新，卖出点也要跟着更新
func maxProfit(prices []int) int {
	n := len(prices)
	maxp, in, out := 0, 10000, 10000
	for i := 0; i < n; i++ {
		if prices[i] < in { // 下降过程中，更新买入点和卖出点
			in = prices[i]
			out = prices[i]
		} else if prices[i] > out { // 上升过程中，更新卖出点
			out = prices[i]
			if out-in > maxp { // 计算更新收益
				maxp = out - in
			}
		}
	}
	return maxp
}

// 优化一
// 消除当天买当天卖的情况
func maxProfit3(prices []int) int {
	n := len(prices)
	in, maxp := 10000, 0
	for i := 0; i < n; i++ {
		if prices[i] < in {
			in = prices[i]
		} else if maxp < prices[i]-in {
			maxp = prices[i] - in
		}
	}
	return maxp
}

// 解法二：买在历史最低点
//
// 因为只能单次买入单次卖出
// 为了获取最大收益只有一条路：低买高卖
// 同时确定最低点和最高点还要满足时间顺序有些太麻烦
// 如果选择在第i天卖出，只要知道前i天的最低点，就能确定最大收益
// 设dp[i]是前i天的最低点，dp[i] = min(dp[i-1],prices[i])
// 最大收益 maxp = max(prices[i]-dp[i])
//
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func maxProfit2(prices []int) int {
	n := len(prices)
	in, maxp := 10000, 0
	for i := 0; i < n; i++ {
		if prices[i] < in {
			in = prices[i]
		}
		if maxp < prices[i]-in {
			maxp = prices[i] - in
		}
	}
	return maxp
}

// 解法一：暴力搜索
// 遍历所有可能的买入点和卖出点
// 假设在第i天买入，第j天卖出，得到的收益为p=prices[j]-prices[i]
// 更新最大收益maxp=max(maxp,p)
//
// 时间复杂度：O(n2)
// 空间复杂度：O(1)
func maxProfit1(prices []int) int {
	n := len(prices)
	maxp := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			p := prices[j] - prices[i]
			if p > maxp {
				maxp = p
			}
		}
	}
	return maxp
}
