package solution

func plusOne(digits []int) []int {
	n := len(digits)
	ans := make([]int, n+1) // 长度+1
	t := 1                  // 加法进位
	for i := n - 1; i >= 0; i-- {
		ans[i+1] = (digits[i] + t) % 10
		t = (digits[i] + t) / 10
	}
	if t != 0 { // 最高位进位
		ans[0] = t
		return ans
	}
	return ans[1:]
}
