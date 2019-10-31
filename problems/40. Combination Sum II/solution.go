package solution

import (
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	ans := make([][]int, 0)
	sort.Ints(candidates)
	backtrack(candidates, target, 0, []int{}, &ans)
	return ans
}

func backtrack(candidates []int, target int, index int, curr []int, ans *[][]int) {
	if target == 0 {
		tmp := make([]int, len(curr))
		copy(tmp, curr)
		*ans = append(*ans, tmp)
		return
	}
	n := len(candidates)
	for i := index; i < n; i++ {
		if i > index && candidates[i] == candidates[i-1] {
			continue
		}
		if candidates[i] <= target {
			curr = append(curr, candidates[i])
			backtrack(candidates, target-candidates[i], i+1, curr, ans)
			curr = curr[:len(curr)-1]
		}
	}
}
