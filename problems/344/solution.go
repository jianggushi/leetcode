package solution

// 双指针，前后两两交换
// 时间复杂度 O(n)
func reverseString(s []byte) {
	n := len(s)
	l, r := 0, n-1
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}

// 借助辅助数组，降低时间复杂度
// 时间复杂度 O(n) 空间复杂度 O(n)
func reverseString2(s []byte) {
	n := len(s)
	ans := make([]byte, n)
	copy(ans, s)
	for i := 0; i < n; i++ {
		s[i] = ans[n-1-i]
	}
}

// 冒泡交换
// 时间复杂度 O(n2)
func reverseString1(s []byte) {
	n := len(s)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			s[j], s[j+1] = s[j+1], s[j]
		}
	}
}
