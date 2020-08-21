package solution

func subarraysDivByK(A []int, K int) int {
	sum := 0
	n := len(A)
	records := make(map[int]int)
	ans := 0
	for i := 0; i < n; i++ {
		sum += A[i]
		tmp := 0
		if sum < 0 {
			tmp = (-sum) % K
		} else {
			tmp = sum % K
		}
		if value, ok := records[tmp]; ok {
			ans += value
			records[tmp]++
		} else {
			records[tmp] = 1
		}
	}
	return ans + records[0]
}
