package solution

import "sort"

// 动态规划
// 因为题目不需要列出所有组合，只需要计算出组合数。因此不需要使用递归来进行搜索。
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if i >= num {
				dp[i] += dp[i-num]
			}
		}
	}
	return dp[target]
}

// backtrack 复杂度高
func combinationSum4_1(nums []int, target int) int {
	ans := 0
	sort.Ints(nums)
	backtrack(nums, target, &ans)
	return ans
}

func backtrack(nums []int, target int, ans *int) {
	if target == 0 {
		*ans++
		return
	}
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] <= target {
			backtrack(nums, target-nums[i], ans)
		}
	}
}
