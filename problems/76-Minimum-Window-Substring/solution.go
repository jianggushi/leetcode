package solution

func minWindow(s string, t string) string {
	ori := map[byte]int{}
	for i := 0; i < len(t); i++ {
		ori[t[i]]++
	}

	check := func(cnt map[byte]int) bool {
		for k, v := range ori {
			if cnt[k] < v {
				return false
			}
		}
		return true
	}

	n := len(s)
	al, ar := -1, n
	find := false
	cnt := map[byte]int{}
	for l, r := 0, 0; r < n; r++ {
		cnt[s[r]]++
		if ori[s[r]] == 0 {
			continue
		}
		for check(cnt) && l <= r {
			find = true
			cnt[s[l]]--
			l++
		}
		if find {
			if r-l+1 < ar-al {
				al, ar = l-1, r
			}
		}
	}
	if al == -1 {
		return ""
	}
	return s[al : ar+1]
}

func minWindow1(s string, t string) string {
	ori := map[byte]int{}
	for i := 0; i < len(t); i++ {
		ori[t[i]]++
	}

	check := func(cnt map[byte]int) bool {
		for k, v := range ori {
			if cnt[k] < v {
				return false
			}
		}
		return true
	}

	n := len(s)
	l, r := 0, n
	find := false
	for i := 0; i < n; i++ {
		cnt := map[byte]int{}
		for j := i; j < n; j++ {
			cnt[s[j]]++
			if check(cnt) && (j-i < r-l) {
				l, r = i, j
				find = true
				break
			}
		}
	}
	if find {
		return s[l : r+1]
	}
	return ""
}
