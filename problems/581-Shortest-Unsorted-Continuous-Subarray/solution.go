package solution

func findUnsortedSubarray(nums []int) int {
	n := len(nums)
	start := 0
	for i := 0; i < n-1; i++ {
		flag := true
		for j := i + 1; j < n; j++ {
			if nums[j] < nums[i] {
				flag = false
			}
		}
		if !flag {
			start = i
			break
		}
	}

	end := 0
	for i := n - 1; i > 0; i-- {
		flag := true
		for j := i - 1; j >= 0; j-- {
			if nums[j] > nums[i] {
				flag = false
			}
		}
		if !flag {
			end = i
			break
		}
	}
	if start == 0 && end == 0 {
		return 0
	}
	return end - start + 1
}
