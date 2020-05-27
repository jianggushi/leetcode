package solution

import "sort"

func findRepeatNumber(nums []int) int {
	n := len(nums)
	for i := 0; i < n; {
		if nums[i] == i {
			i++
			continue
		}
		if nums[i] == nums[nums[i]] {
			return nums[i]
		}
		nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
	}
	return -1
}

// 哈希表辅助检测
func findRepeatNumber2(nums []int) int {
	m := make(map[int]struct{})
	n := len(nums)
	for i := 0; i < n; i++ {
		if _, ok := m[nums[i]]; ok {
			return nums[i]
		} else {
			m[nums[i]] = struct{}{}
		}
	}
	return -1
}

// 先排序再检测
func findRepeatNumber1(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n-1; i++ {
		if nums[i] == nums[i+1] {
			return nums[i]
		}
	}
	return -1
}
