package solution

func combine(n int, k int) [][]int {
	ans := make([][]int, 0)
	backtrack(n, k, 1, []int{}, &ans)
	return ans
}

func backtrack(n int, k int, first int, curr []int, ans *[][]int) {
	if len(curr) == k {
		// 复制 curr
		tmp := make([]int, k)
		copy(tmp, curr)
		*ans = append(*ans, tmp)
	}
	for i := first; i <= n; i++ {
		curr = append(curr, i)
		backtrack(n, k, i+1, curr, ans)
		curr = curr[:len(curr)-1]
	}
}
