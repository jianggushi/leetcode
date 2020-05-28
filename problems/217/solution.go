package solution

// 哈希表
func containsDuplicate(nums []int) bool {
	n := len(nums)
	mmap := make(map[int]int)
	for i := 0; i < n; i++ {
		if mmap[nums[i]] > 0 {
			return true
		}
		mmap[nums[i]]++
	}
	return false
}

// 暴力搜索
func containsDuplicate1(nums []int) bool {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if nums[j] == nums[i] {
				return true
			}
		}
	}
	return false
}
