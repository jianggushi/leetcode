package solution

func searchRange(nums []int, target int) []int {
	l := searchLeft(nums, target)
	r := searchRight(nums, target)
	return []int{l, r}
}

func searchLeft(nums []int, target int) int {
	n := len(nums)
	l, r := 0, n-1
	for l <= r {
		m := (l + r) / 2
		if nums[m] < target {
			l = m + 1
		} else if nums[m] > target {
			r = m - 1
		} else {
			if m == 0 || nums[m-1] != target {
				return m
			}
			r = m - 1
		}
	}
	return -1
}

func searchRight(nums []int, target int) int {
	n := len(nums)
	l, r := 0, n-1
	for l <= r {
		m := (l + r) / 2
		if nums[m] < target {
			l = m + 1
		} else if nums[m] > target {
			r = m - 1
		} else {
			if m == n-1 || nums[m+1] != target {
				return m
			}
			l = m + 1
		}
	}
	return -1
}
