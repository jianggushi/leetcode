package solution

// 解法一
// 从后向前查找，如果当前数小于后面一个数，说明可以替换当前数
// 然后，从当前数后面开始的数中（前面一步已经保证后面的数是倒序的），从后向前查找，找到第一个比当前数大的一个数，肯定可以找到
// 和这个数交换，那么后面的数还是倒序的
// 然后再把后面的数进行排序，从小到大，可以通过首尾交换做到
func nextPermutation(nums []int) {
	n := len(nums)
	l, r := 0, n-1
	for i := n - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			for j := n - 1; j > i; j-- {
				if nums[j] > nums[i] {
					nums[i], nums[j] = nums[j], nums[i]
					break
				}
			}
			l = i + 1
			break
		}
	}
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}
