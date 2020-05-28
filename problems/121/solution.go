package solution

import "math"

func maxProfit(prices []int) int {
	n := len(prices)
	ans := 0
	in, out := math.MaxInt64, math.MaxInt64
	for i := 0; i < n; i++ {
		if prices[i] < in { // 下降过程中，更新买入点和卖出点
			in = prices[i]
			out = prices[i]
		} else if prices[i] > out { // 上升过程中，更新卖出点
			out = prices[i]
			if out-in > ans { // 计算更新收益
				ans = out - in
			}
		}
	}
	return ans
}

// 优化二
// 时间复杂度 O(n) 空间复杂度 O(1)
// 优化一种确定了买入点，在每一个位置都计算卖出的收益，但有时候的计算是无意义的
// 比如在下降阶段，买入点一直在更新，收益在无意义的更新检查
// 为了避免检查一些无意义的卖出点，要对之前的卖出点进行记录，只有当前价格大于之前的卖出点再更新卖出点和收益
// ! 在下降阶段，买入点一直在更新，卖出点也要跟着更新
func maxProfit3(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	ans := 0
	in, out := prices[0], prices[0] // 初始买入点和卖出点
	for i := 1; i < n; i++ {
		if prices[i] > out { // 如果当前价格大于卖出点，更新卖出点
			out = prices[i]
			if out-in > ans { // 计算更新收益
				ans = out - in
			}
		}
		if prices[i] < in { // 如果当前价格小于买入点，更新买入点和卖出点
			in = prices[i]
			out = prices[i]
		}
	}
	return ans
}

func maxProfit21(prices []int) int {
	n := len(prices)
	ans, in := 0, math.MaxInt64
	for i := 0; i < n; i++ {
		if prices[i] < in { // 如果当前价格小于买入点，更新买入点
			in = prices[i]
		} else if prices[i]-in > ans {
			ans = prices[i] - in // 更新收益
		}
	}
	return ans
}

// 优化一
// 时间复杂度 O(n) 空间复杂度 O(1)
// 思路是确定买入点，要买入在低点
func maxProfit2(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	ans := 0
	in := prices[0] // 初始买入点
	for i := 1; i < n; i++ {
		tmp := prices[i] - in // 计算更新收益
		if tmp > ans {
			ans = tmp
		}
		if prices[i] < in { // 如果当前价格小于买入点，更新买入点
			in = prices[i]
		}
	}
	return ans
}

// 暴力搜索
// 时间复杂度 O(n2) 空间复杂度 O(1)
// 遍历所有可能的买入点和卖出点
func maxProfit1(prices []int) int {
	n := len(prices)
	ans := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			tmp := prices[j] - prices[i]
			if tmp > ans {
				ans = tmp
			}
		}
	}
	return ans
}
