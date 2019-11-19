package solution

import "sort"

// 弗洛伊德的乌龟和兔子（循环检测）
func findDuplicate(nums []int) int {
	slow, fast := nums[nums[0]], nums[nums[nums[0]]]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	p1, p2 := nums[0], slow
	for p1 != p2 {
		p1 = nums[p1]
		p2 = nums[p2]
	}
	return p1
}

// 使用 map 去重
// 时间复杂度 O(n), 空间复杂度 O(n)
func findDuplicate2(nums []int) int {
	n := len(nums)
	mmap := make(map[int]bool, n)

	for i := 0; i < n; i++ {
		if _, exist := mmap[nums[i]]; exist {
			return nums[i]
		}
		mmap[nums[i]] = true
	}
	return 0
}

// 先排序
// 时间复杂度 O(nlogn)
func findDuplicate1(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	for i := 1; i < n; i++ {
		if nums[i] == nums[i-1] {
			return nums[i]
		}
	}
	return 0
}
