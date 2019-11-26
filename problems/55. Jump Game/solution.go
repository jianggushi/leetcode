package solution

// 解法二
// 从后向前尝试
func canJump(nums []int) bool {
	n := len(nums)
	pos := n - 1
	for i := n - 1; i >= 0; i-- {
		if i+nums[i] >= pos {
			pos = i
		}
	}
	return pos == 0
}

// 解法一
// 回溯算法
func canJump1(nums []int) bool {
	ans := false
	backtrack(nums, 0, &ans)
	return ans
}

func backtrack(nums []int, pos int, ans *bool) {
	n := len(nums)
	if pos == n-1 {
		*ans = true
		return
	}
	for i := 1; i <= nums[pos] && pos+i < n; i++ {
		backtrack(nums, pos+i, ans)
		// 遍历过的节点置为 -1，避免重复遍历
		nums[pos+i] = -1
	}
}
