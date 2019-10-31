package solution

func subsets(nums []int) [][]int {
	n := len(nums)
	ans := make([][]int, 1)
	ans[0] = []int{}
	for i := 0; i < n; i++ {
		m := len(ans)
		for j := 0; j < m; j++ {
			tmp := make([]int, len(ans[j])+1)
			copy(tmp, ans[j])
			tmp[len(tmp)-1] = nums[i]
			ans = append(ans, tmp)
		}
	}
	return ans
}
