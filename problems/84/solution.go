package solution

import (
	"math"
)

// 递归求解
// 假设 f(i) 是前 i 个元素组成的最大矩形
// 如果知道了 f(i-1) 的值，怎么推导得到 f(i)
// f(i) = max[f(i-1), 0到i的最小元素*(i+1), 以i为最小的连续元素个数*间距]
func largestRectangleArea3(heights []int) int {
	n := len(heights)
	dp := 0
	for i := 0; i < n; i++ {
		t, res := heights[i], heights[i]
		for j := i - 1; j >= 0; j-- {
			if heights[j] < t {
				t = heights[j]
			}
			if t*(i-j+1) > res {
				res = t * (i - j + 1)
			}
		}
		if res > dp {
			dp = res
		}
	}
	return dp
}

// 递归求解
// 假设 f(i) 是前 i 个元素组成的最大矩形
// 如果知道了 f(i-1) 的值，怎么推导得到 f(i)
// f(i) = max[f(i-1), 0到i的最小元素*(i+1), 以i为最小的连续元素个数*间距]
func largestRectangleArea2(heights []int) int {
	n := len(heights)
	if n == 0 {
		return 0
	}
	a1 := largestRectangleArea2(heights[:n-1])
	a2 := heights[n-1]
	t := heights[n-1]
	for i := n - 2; i >= 0; i-- {
		if heights[i] < t {
			t = heights[i]
		}
		if t*(n-i) > a2 {
			a2 = t * (n - i)
		}
	}
	if a1 > a2 {
		return a1
	}
	return a2
}

// 暴力求解
// 遍历所有可能
func largestRectangleArea1(heights []int) int {
	n := len(heights)
	ans := 0
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			tmp := min(heights[i:j+1]) * (j - i + 1)
			if tmp > ans {
				ans = tmp
			}
		}
	}
	return ans
}

func min(heights []int) int {
	n := len(heights)
	ans := math.MaxInt64
	for i := 0; i < n; i++ {
		if heights[i] < ans {
			ans = heights[i]
		}
	}
	return ans
}
