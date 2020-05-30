package solution

// 优化
// 因为要求的是连续子序列
func findLengthOfLCIS(nums []int) int {
	n := len(nums)
	ans := 0
	i, j := 0, 0
	for i < n {
		tmp := 1
		for j = i + 1; j < n; j++ {
			if nums[j] <= nums[j-1] {
				break
			}
			tmp++
		}
		if tmp > ans {
			ans = tmp
		}
		i = j // 直接跳过前面递增的子序列，从后面查找
	}
	return ans
}

// 暴力搜索
func findLengthOfLCIS1(nums []int) int {
	n := len(nums)
	ans := 0
	for i := 0; i < n; i++ {
		tmp := 1
		for j := i + 1; j < n; j++ {
			if nums[j] <= nums[j-1] {
				break
			}
			tmp++
		}
		if tmp > ans {
			ans = tmp
		}
	}
	return ans
}
