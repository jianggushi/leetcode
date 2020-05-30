package solution

// 使用 map 辅助计数
func isAnagram(s string, t string) bool {
	smap := make(map[rune]int)
	for _, r := range s {
		smap[r]++
	}
	tmap := make(map[rune]int)
	for _, r := range t {
		tmap[r]++
	}
	if len(smap) != len(tmap) {
		return false
	}
	for r, v1 := range smap {
		v2, ok := tmap[r]
		if !ok || v1 != v2 {
			return false
		}
	}
	return true
}
