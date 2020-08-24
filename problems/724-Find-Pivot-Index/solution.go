package solution

func pivotIndex(nums []int) int {
	n := len(nums)
	lsum, rsum := 0, 0
	for i := 0; i < n; i++ {
		rsum += nums[i]
	}

	for p := 0; p < n; p++ {
		if p > 0 {
			lsum += nums[p-1]
		}
		rsum -= nums[p]
		if lsum == rsum {
			return p
		}
	}
	return -1
}
