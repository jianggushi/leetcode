package solution

import "sort"

func groupAnagrams(strs []string) [][]string {
	mmap := make(map[[26]byte][]string)
	n := len(strs)
	for i := 0; i < n; i++ {
		h := hash(strs[i])
		mmap[h] = append(mmap[h], strs[i])
	}
	ans := [][]string{}
	for _, vals := range mmap {
		ans = append(ans, vals)
	}
	return ans
}
func hash(s string) [26]byte {
	ans := [26]byte{}
	n := len(s)
	for i := 0; i < n; i++ {
		ans[s[i]-'a']++
	}
	return ans
}

type byteSlice []byte

func (p byteSlice) Len() int           { return len(p) }
func (p byteSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p byteSlice) Less(i, j int) bool { return p[i] < p[j] }

// 解法一
// 排序归类
func groupAnagrams1(strs []string) [][]string {
	mmap := make(map[string][]string)
	n := len(strs)
	for i := 0; i < n; i++ {
		t := byteSlice(strs[i])
		sort.Sort(t)
		t1 := string(t)
		mmap[t1] = append(mmap[t1], strs[i])
	}
	ans := [][]string{}
	for _, vals := range mmap {
		ans = append(ans, vals)
	}
	return ans
}
