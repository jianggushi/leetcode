package solution

func countSubstrings(s string) int {
	n := len(s)
	ans := 0
	for i := 0; i < n; i++ {
		ans = 1 + ans
		for j := i - 1; j >= 0; j-- {
			if palindromic(s[j : i+1]) {
				ans++
			}
		}
	}
	return ans
}

func countSubstrings1(s string) int {
	n := len(s)
	ans := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			if palindromic(s[i:j]) {
				ans++
			}
		}
	}
	return ans
}

func palindromic(s string) bool {
	i := 0
	j := len(s) - 1
	for i <= j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
