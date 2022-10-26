// Source : https://leetcode.cn/problems/running-sum-of-1d-array/
// Author : jianggushi
// Date   : 2022-10-26

package solution

// 解法二：原地修改
// 原地修改输入数组 nums，降低空间复杂度
// 根据得到的递推公式
// runningSum[i] = runningSum[i-1] + nums[i]
// 当 i = 0 时, runningSum[0] = nums[0]
// 可以让 nums[i] = nums[i-1] + nums[i]
func runningSum(nums []int) []int {
	n := len(nums)
	for i := 1; i < n; i++ {
		nums[i] += nums[i-1]
	}
	return nums
}

// 解法一：遍历数组滚动求和
// 根据题目描述
// runningSum[i] = nums[0]+nums[1]+...+nums[i]
// runningSum[i-1] = nums[0]+nums[1]+...+nums[i-1]
// 可以推导出递推公式
// runningSum[i] = runningSum[i-1] + nums[i]
// 当 i = 0 时, runningSum[0] = nums[0]
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func runningSum1(nums []int) []int {
	n := len(nums)
	ans := make([]int, len(nums))
	ans[0] = nums[0]
	for i := 1; i < n; i++ {
		ans[i] = ans[i-1] + nums[i]
	}
	return ans
}
