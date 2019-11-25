package solution

// 解法二
// 对解法一进行总结归纳到一个函数中
func search(nums []int, target int) int {
	n := len(nums)
	l, r := 0, n-1
	for l <= r {
		m := (l + r) / 2
		if nums[m] < target {
			if nums[m] < nums[l] && target >= nums[l] {
				r = m - 1
			} else {
				l = m + 1
			}
		} else if nums[m] > target {
			if nums[m] > nums[r] && target <= nums[r] {
				l = m + 1
			} else {
				r = m - 1
			}
		} else {
			return m
		}
	}
	return -1
}

// 解法一
// 分情况进行分析解决
func search1(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	l, r := 0, n-1
	ans := -1
	// 存在翻转
	if nums[l] > nums[r] {
		if target >= nums[l] {
			// target 在左半侧
			ans = leftBSearch(nums, target)
		} else if target <= nums[r] {
			// target 在右半侧
			ans = rightBSearch(nums, target)
		}
	} else {
		// 顺序不存在翻转，二分查找
		ans = bSearch(nums, target)
	}
	return ans
}

// target 在左半侧，进行二分查找
// nums[m] < target 有两种情况
// 1. nums[m] 落在左半侧，nums[m] >= nums[l]，只需要 l = m+1
// 2. nums[m] 落在右半侧，nums[m] < nums[l]，需要 r = m-1
// nums[m] > target 此时 nums[m] 肯定落在左半侧，只需要 r = m-1
func leftBSearch(nums []int, target int) int {
	n := len(nums)
	l, r := 0, n-1
	for l <= r {
		m := (l + r) / 2
		if nums[m] < target {
			if nums[m] >= nums[l] {
				l = m + 1
			} else {
				r = m - 1
			}
		} else if nums[m] > target {
			r = m - 1
		} else {
			return m
		}
	}
	return -1
}

// target 在右半侧，进行二分查找
// nums[m] < target 此时 nums[m] 肯定落在右半侧，只需要 l = m+1
// nums[m] > target 有两种情况
// 1. nums[m] 落在左半侧，nums[m] > nums[r]，需要 l = m+1
// 2. nums[m] 落在右半侧，nums[m] <= nums[r]，需要 r = m-1

func rightBSearch(nums []int, target int) int {
	n := len(nums)
	l, r := 0, n-1
	for l <= r {
		m := (l + r) / 2
		if nums[m] < target {
			l = m + 1
		} else if nums[m] > target {
			if nums[m] > nums[r] {
				l = m + 1
			} else {
				r = m - 1
			}
		} else {
			return m
		}
	}
	return -1
}

// nums 有序，进行标准的二分查找
func bSearch(nums []int, target int) int {
	n := len(nums)
	l, r := 0, n-1
	for l <= r {
		m := (l + r) / 2
		if nums[m] < target {
			l = m + 1
		} else if nums[m] > target {
			r = m - 1
		} else {
			return m
		}
	}
	return -1
}
