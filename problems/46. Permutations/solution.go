package solution

func permute(nums []int) [][]int {
	ans := make([][]int, 0)
	backtrack(nums, 0, &ans)
	return ans
}

func backtrack(nums []int, first int, ans *[][]int) {
	if first == len(nums) {
		tmp := make([]int, first)
		copy(tmp, nums)
		*ans = append(*ans, tmp)
		return
	}
	for i := first; i < len(nums); i++ {
		nums[first], nums[i] = nums[i], nums[first]
		backtrack(nums, first+1, ans)
		nums[first], nums[i] = nums[i], nums[first]
	}
}
