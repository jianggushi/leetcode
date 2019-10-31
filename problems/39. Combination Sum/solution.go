package solution

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	ans := make([][]int, 0)
	sort.Ints(candidates)
	backtrack(candidates, target, 0, []int{}, &ans)
	return ans
}

func backtrack(candidates []int, target int, first int, curr []int, ans *[][]int) {
	if target == 0 {
		tmp := make([]int, len(curr))
		copy(tmp, curr)
		*ans = append(*ans, tmp)
		return
	}
	n := len(candidates)
	for i := first; i < n; i++ {
		if candidates[i] <= target {
			curr = append(curr, candidates[i])
			backtrack(candidates, target-candidates[i], i, curr, ans)
			curr = curr[:len(curr)-1]
		}
	}
}
