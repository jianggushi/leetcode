package solution

// 复用结果数组，降低空间复杂度
func productExceptSelf(nums []int) []int {
	n := len(nums)

	ans := make([]int, n)
	ans[0] = 1
	for i := 1; i < n; i++ {
		ans[i] = ans[i-1] * nums[i-1]
	}
	r := 1
	for i := n - 2; i >= 0; i-- {
		r *= nums[i+1]
		ans[i] *= r
	}
	return ans
}

// 降低时间复杂度 O(n)
func productExceptSelf2(nums []int) []int {
	n := len(nums)

	l := make([]int, n)
	r := make([]int, n)
	l[0] = 1
	for i := 1; i < n; i++ {
		l[i] = l[i-1] * nums[i-1]
	}
	r[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		r[i] = r[i+1] * nums[i+1]
	}
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = l[i] * r[i]
	}
	return ans
}

// 暴力求解
// 时间复杂度 O(n2)
// 空间复杂度 O(n)，计算了结果数组的空间
func productExceptSelf1(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = 1
		for j := 0; j < n; j++ {
			if j != i {
				ans[i] *= nums[j]
			}
		}
	}
	return ans
}
