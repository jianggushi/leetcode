package solution

// 借助 map 统计
func firstUniqChar(s string) int {
	n := len(s)
	mmap := make(map[byte]int)
	for i := 0; i < n; i++ {
		mmap[s[i]]++
	}
	for i := 0; i < n; i++ {
		if mmap[s[i]] == 1 {
			return i
		}
	}
	return -1
}
