package solution

// 先遍历 candies 找到最大值 maxCandies
// 然后再遍历 candies 每个元素加上 extraCandies
// 如果大于或等于 maxCandies 那么就是候选设置为 true
func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxCandies := 0
	n := len(candies)
	for i := 0; i < n; i++ {
		if candies[i] > maxCandies {
			maxCandies = candies[i]
		}
	}
	ans := make([]bool, n)
	for i := 0; i < n; i++ {
		if candies[i]+extraCandies >= maxCandies {
			ans[i] = true
		}
	}
	return ans
}
