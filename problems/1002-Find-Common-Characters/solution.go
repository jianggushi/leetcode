package solution

func commonChars(A []string) []string {
	ans := []string{}
	n := len(A)
	if n == 0 {
		return ans
	}
	count := make([]int, len(A[0]))
	for i := 1; i < n; i++ {
		for _, c := range A[i] {
			for j, c2 := range A[0] {
				if c == c2 && count[j] == i-1 {
					count[j]++
					break
				}
			}
		}
	}
	for i, c := range A[0] {
		if count[i] == n-1 {
			ans = append(ans, string(c))
		}
	}
	return ans
}
