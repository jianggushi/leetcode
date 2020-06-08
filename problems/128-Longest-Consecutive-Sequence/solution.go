package solution

import (
	"sort"
)

// 使用哈希表存储元素，快速查找
// 时间复杂度 O(n2)
func longestConsecutive(nums []int) int {
	n := len(nums)
	mmap := make(map[int]int)
	for i := 0; i < n; i++ {
		mmap[nums[i]]++
	}
	ans, curr := 0, 0
	for i := range mmap {
		curr = 1
		// 如果 i-1 在数组中，那么可以跳过 i
		if mmap[i-1] > 0 {
			continue
		}
		for j := i + 1; ; j++ {
			if mmap[j] > 0 {
				curr++
			} else {
				break
			}
		}
		if curr > ans {
			ans = curr
		}
	}
	return ans
}

// 使用哈希表存储元素，快速查找
// 时间复杂度 O(n2)
func longestConsecutive2(nums []int) int {
	n := len(nums)
	mmap := make(map[int]int)
	for i := 0; i < n; i++ {
		mmap[nums[i]]++
	}
	ans, curr := 0, 0
	for i := range mmap {
		curr = 1
		for j := i + 1; ; j++ {
			if mmap[j] > 0 {
				curr++
			} else {
				break
			}
		}
		if curr > ans {
			ans = curr
		}
	}
	return ans
}

// 先排序再遍历检查
// 注意数组中可能有重复元素
// 排序 O(nlogn)
// 遍历 O(n)
func longestConsecutive1(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	sort.Ints(nums)
	ans, curr := 1, 1
	for i := 1; i < n; i++ {
		if nums[i] == nums[i-1] {
			continue
		} else if nums[i]-nums[i-1] == 1 {
			curr++
		} else {
			if curr > ans {
				ans = curr
			}
			curr = 1
		}

	}
	if curr > ans {
		ans = curr
	}
	return ans
}
