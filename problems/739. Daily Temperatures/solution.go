package solution

// 从右向左遍历
func dailyTemperatures(T []int) []int {
	n := len(T)
	ans := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j += ans[j] {
			if T[j] > T[i] {
				ans[i] = j - i
				break
			} else if ans[j] == 0 {
				break
			}
		}
	}
	return ans
}

// 从左向右遍历
func dailyTemperatures1(T []int) []int {
	n := len(T)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if T[j] > T[i] {
				ans[i] = j - i
				break
			}
		}
	}
	return ans
}
