package solution

// 回溯 - 递归
func subsets(nums []int) [][]int {
	ans := make([][]int, 0)
	backtrack(nums, 0, []int{}, &ans)
	return ans
}

func backtrack(nums []int, first int, curr []int, ans *[][]int) {
	tmp := make([]int, len(curr))
	copy(tmp, curr)
	*ans = append(*ans, tmp)
	for i := first; i < len(nums); i++ {
		curr = append(curr, nums[i])
		backtrack(nums, i+1, curr, ans)
		curr = curr[:len(curr)-1]
	}
}

// 迭代
func subsets1(nums []int) [][]int {
	n := len(nums)
	// 结果列表
	ans := make([][]int, 1)
	ans[0] = []int{}
	// 遍历每一个数字
	for i := 0; i < n; i++ {
		m := len(ans)
		// 和结果列表中的组合，形成新的结果并加到结果列表
		for j := 0; j < m; j++ {
			tmp := make([]int, len(ans[j])+1)
			copy(tmp, ans[j])
			tmp[len(tmp)-1] = nums[i]
			ans = append(ans, tmp)
		}
	}
	return ans
}
