package solution

func combinationSum(candidates []int, target int) [][]int {
	ans := make([][]int, 0)
	result := make([]int, 0)
	combination(candidates, target, 0, result, &ans)
	return ans
}

func combination(candidates []int, target int, index int, result []int, ans *[][]int) {
	if target == 0 {
		*ans = append(*ans, result)
		return
	}
	n := len(candidates)
	for i := index; i < n; i++ {
		if candidates[i] <= target {
			tmp := make([]int, len(result)+1)
			copy(tmp, result)
			tmp[len(result)] = candidates[i]
			combination(candidates, target-candidates[i], i, tmp, ans)
		}
	}
}
