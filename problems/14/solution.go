package solution

func longestCommonPrefix(strs []string) string {
	n := len(strs)
	if n == 0 {
		return ""
	}
	if n == 1 {
		return strs[0]
	}

	i := 0
	for ; i < len(strs[0]); i++ {
		j := 1
		for ; j < n; j++ {
			if i >= len(strs[j]) {
				break
			}
			if strs[j][i] != strs[0][i] {
				break
			}
		}
		if j != n {
			break
		}
	}
	return strs[0][0:i]
}
