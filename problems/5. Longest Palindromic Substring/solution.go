package solution

// 解法三
// 从中心向两边扩展
func longestPalindrome(s string) string {
	n := len(s)
	ans := ""
	for i := 0; i < n; i++ {
		s1 := expandAroundCenter(s, i, i)
		if len(s1) > len(ans) {
			ans = s1
		}
		s2 := expandAroundCenter(s, i, i+1)
		if len(s2) > len(ans) {
			ans = s2
		}

	}
	return ans
}

func expandAroundCenter(s string, l int, r int) string {
	n := len(s)
	for l >= 0 && r < n && s[l] == s[r] {
		l--
		r++
	}
	return s[l+1 : r]
}

// 解法二
// 从最长向最短逼近
func longestPalindrome2(s string) string {
	n := len(s)
	l, r := 0, n
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			if isPalindrome(s[l+j : r-i+j]) {
				return s[l+j : r-i+j]
			}
		}
	}
	return ""
}

// 解法一
// 暴力搜索法，遍历所有字符串组合方式，然后判断是否为回文串
func longestPalindrome1(s string) string {
	n := len(s)
	ans := ""
	for i := 0; i < n; i++ {
		for j := i + 1; j < n+1; j++ {
			if isPalindrome(s[i:j]) && len(s[i:j]) > len(ans) {
				ans = s[i:j]
			}
		}
	}
	return ans
}

func isPalindrome(s string) bool {
	n := len(s)
	i, j := 0, n-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
