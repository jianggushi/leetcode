package solution

// 遍历+交换
func moveZeroes(nums []int) {
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			j := i + 1 // 找到下一个非零元素
			for j < n && nums[j] == 0 {
				j++
			}
			if j < n { // 找到交换
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}
