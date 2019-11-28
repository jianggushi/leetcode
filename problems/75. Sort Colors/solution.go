package solution

import "sort"

// 解法二
// 因为只有 0, 1, 2 三种数字，使用双指针从头尾进行遍历
// 以右侧为准
// (1) 如果 nums[r] == 0，
// 那么需要从左边找到一个不为 0 的数字进行交换，找到的同时 l 向右移动，
// 交换后，如果 nums[r] == 0 或 2，r 需要向左移动
// (2) 如果 nums[r] == 1，
// 如果左侧 nums[l] == 0，l 向右移动
// 如果左侧 nums[l] == 1，那么需要从右侧开始向前找 2，如果找到和 nums[r] 交换，之后 r 需要向左移动
// 如果左侧 nums[l] == 2，左右交换，同时 r 向左移动
// (3) 如果 nums[r] == 2，
// 只需要 r 向左移动
func sortColors(nums []int) {
	l, r := 0, len(nums)-1
	for l < r {
		if nums[r] == 0 {
			for i := l; i < r; i++ {
				if nums[i] > 0 {
					nums[i], nums[r] = nums[r], nums[i]
					l = i + 1
					break
				}
			}
			if nums[r] == 2 || nums[r] == 0 {
				r--
			}
		} else if nums[r] == 1 {
			if nums[l] == 0 {
				l++
			} else if nums[l] == 1 {
				for i := r; i > l; i-- {
					if nums[i] == 2 {
						nums[i], nums[r] = nums[r], nums[i]
						break
					}
				}
				r--
			} else {
				nums[l], nums[r] = nums[r], nums[l]
				r--
			}
		} else {
			r--
		}
	}
}

// 解法一
// 排序（库函数）
func sortColors1(nums []int) {
	sort.Ints(nums)
}
