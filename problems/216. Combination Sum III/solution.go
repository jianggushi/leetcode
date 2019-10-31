package solution

func combinationSum3(k int, n int) [][]int {
	ans := make([][]int, 0)
	backtrack(k, n, 1, []int{}, &ans)
	return ans
}

func backtrack(k int, n int, first int, curr []int, ans *[][]int) {
	if n == 0 && len(curr) == k {
		tmp := make([]int, k)
		copy(tmp, curr)
		*ans = append(*ans, tmp)
		return
	}
	if n < 0 {
		return
	}
	if len(curr) >= k {
		return
	}
	for i := first; i <= 9; i++ {
		curr = append(curr, i)
		backtrack(k, n-i, i+1, curr, ans)
		curr = curr[:len(curr)-1]
	}
}
