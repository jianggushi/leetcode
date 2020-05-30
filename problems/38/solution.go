package solution

import (
	"fmt"
)

// 递归求解
func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	ans := ""
	s := countAndSay(n - 1)
	m := len(s)
	for i := 0; i < m; {
		j := i + 1
		for j < m && s[j] == s[i] {
			j++
		}
		ans += fmt.Sprintf("%d%c", j-i, s[i])
		i = j
	}
	return ans
}
