package solution

import (
	"strings"
)

func reverseWords(s string) string {
	n := len(s)
	i := n - 1
	ans := []string{}
	for i >= 0 {
		// trim tail space
		for i >= 0 && s[i] == ' ' {
			i--
		}
		if i == -1 {
			break
		}
		j := i
		for j >= 0 && s[j] != ' ' {
			j--
		}
		ans = append(ans, s[j+1:i+1])
		i = j
	}
	return strings.Join(ans, " ")
}
