package solution

func removeDuplicates(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		j := i + 1
		for ; j < n; j++ {
			if nums[j] != nums[i] {
				break
			}
		}
		if j > i+1 {
			k := i + 1
			for j < n {
				nums[k] = nums[j]
				k++
				j++
			}
			n = k
		}
	}
	return n
}
