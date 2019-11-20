package solution

// 解法二
// 基于解法一的思想，改进代码组织方式，在一次循环中进行计算两种情况
func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	dp00, dp01 := 0, 0
	dp10, dp11 := 0, 0
	for i := 1; i < n; i++ {
		dp00, dp01 = dp01, max(dp01, dp00+nums[i])
		dp10, dp11 = dp11, max(dp11, dp10+nums[i-1])
	}
	return max(dp01, dp11)
}

// 解法一
// 相比 House Robber, 这里的房子是环装, 首尾相接的, 这意味着第一个房子和最后一个房子只能选择偷一个
// 如果不偷第一个房子, 那么问题简化为 rob(nums[1:])
// 如果不偷最后一个房子，那么问题简化为 rob(nums[:n-1])
// 最终的结果就是 max(rob(nums[1:]), rob(nums[:n-1]))
func rob1(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	return max(robp(nums[1:]), robp(nums[:n-1]))
}

func robp(nums []int) int {
	n := len(nums)
	dp0, dp1 := 0, 0
	for i := 0; i < n; i++ {
		dp0, dp1 = dp1, max(dp1, dp0+nums[i])
	}
	return dp1
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
