package solution

import "math"

func maxSubArray1(nums []int) int {
	n := len(nums)
	result := nums[0]
	for i := 0; i < n; i++ {
		for j := i + 1; j < n+1; j++ {
			tmp := sum(nums[i:j])
			if tmp > result {
				result = tmp
			}
		}
	}
	return result
}

func sum(nums []int) int {
	n := len(nums)
	total := 0
	for i := 0; i < n; i++ {
		total += nums[i]
	}
	return total
}

// 暴力解法
// 遍历所有可能的连续子数组
func maxSubArray(nums []int) int {
	n := len(nums)
	ans := math.MinInt64
	for i := 0; i < n; i++ {
		tmp := 0
		for j := i; j < n; j++ {
			tmp += nums[j]
			if tmp > ans {
				ans = tmp
			}
		}
	}
	return ans
}
