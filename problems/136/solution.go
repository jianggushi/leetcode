package solution

// 位运算，异或
func singleNumber(nums []int) int {
	n := len(nums)
	ans := 0
	for i := 0; i < n; i++ {
		ans ^= nums[i]
	}
	return ans
}

// 哈希表
func singleNumber1(nums []int) int {
	n := len(nums)
	mmap := make(map[int]int)
	for i := 0; i < n; i++ {
		mmap[nums[i]]++
	}
	for k, v := range mmap {
		if v == 1 {
			return k
		}
	}
	return -1
}
