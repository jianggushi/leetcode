package solution

// 使用额外的数组存储替换结果
// 时间复杂度 O(n)
// 空间复杂度 O(n)
func replaceSpace(s string) string {
	bs := make([]rune, 0, len(s))
	for _, v := range s {
		if v == ' ' {
			bs = append(bs, '%', '2', '0')
		} else {
			bs = append(bs, v)
		}
	}
	return string(bs)
}
