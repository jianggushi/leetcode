package solution

// 哈希表辅助计数
func intersect(nums1 []int, nums2 []int) []int {
	mmap := make(map[int]int)
	n1 := len(nums1)
	for i := 0; i < n1; i++ {
		mmap[nums1[i]]++
	}
	n2 := len(nums2)
	ans := []int{}
	for i := 0; i < n2; i++ {
		if mmap[nums2[i]] > 0 {
			ans = append(ans, nums2[i])
			mmap[nums2[i]]--
		}
	}
	return ans
}
