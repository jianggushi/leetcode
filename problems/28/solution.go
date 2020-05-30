package solution

// 暴力搜索
func strStr1(haystack string, needle string) int {
	n := len(haystack)
	m := len(needle)
	if m == 0 {
		return 0
	}
	for i := 0; i < n; i++ {
		j := i
		for ; j < n && j < m+i; j++ {
			if haystack[j] != needle[j-i] {
				break
			}
		}
		if j == m+i {
			return i
		}
	}
	return -1
}
