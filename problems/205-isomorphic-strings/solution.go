// Source : https://leetcode.cn/problems/isomorphic-strings/
// Author : jianggushi
// Date   : 2022-10-27

// 分析
//
// 每个s中的字符都映射到一个t中的字符
// 不同字符不能映射到同一个字符
// 相同字符只能映射到同一个字符
// 字符可以映射到自己本身
// 其实就是要求s和t中的字符满足一一映射，考虑使用哈希表存储这种映射关系，方便查找

package solution

// 解法二：两个哈希表
//
// 只用一个哈希表记录s->t的映射关系，在检测t->s的映射关系时，需要遍历哈希表，时间复杂度高
// 可以同时使用另一个哈希表记录t->s的映射关系，空间换时间
// 从左向右遍历两个字符串，检测新的映射和原来的映射是否存在冲突
// 如果s[i]已经存在映射且映射的字符和t[i]不同，则存在冲突
// 如果t[i]已经存在映射且映射的字符和s[i]不同，则存在冲突
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func isIsomorphic(s string, t string) bool {
	n := len(s)
	s2t := make(map[byte]byte, 0)
	t2s := make(map[byte]byte, 0)
	for i := 0; i < n; i++ {
		if ti, ok := s2t[s[i]]; ok {
			if ti != t[i] {
				return false
			}
		} else {
			if si, ok := t2s[t[i]]; ok {
				if si != s[i] {
					return false
				}
			}
			s2t[s[i]] = t[i]
			t2s[t[i]] = s[i]
		}
	}
	return true
}

// 解法一：哈希表
//
// 为了记录s和t之间字符的一一映射关系，需要用到哈希表
// 哈希表记录从s到t的映射关系
// 举例：s=bar, t=foo
// s到t的映射关系：b->f, a->o, r->o
// 每个字符都有一个映射关系，但a和r映射到了同一个字符o
// 如果要检测是否有字符映射到同一个字符，需要遍历整个哈希表，时间复杂度比较高
// 从左向右遍历两个字符串，检测新的映射和原来的映射是否存在冲突
// 如果s[i]已经存在映射且映射的字符和t[i]不同，则存在冲突
// 如果t[i]被映射过且原字符和s[i]不同，则存在冲突，需要遍历哈希表
// 时间复杂度：O(n2)
// 空间复杂度：O(n)
func isIsomorphic1(s string, t string) bool {
	n := len(s)
	s2t := make(map[byte]byte, 0)
	for i := 0; i < n; i++ {
		if ti, ok := s2t[s[i]]; ok {
			if ti != t[i] { // 检测s->t的映射关系
				return false
			}
		} else {
			for si, ti := range s2t {
				if ti == t[i] && si != s[i] { // 检测t->s的映射关系
					return false
				}
			}
			s2t[s[i]] = t[i]
		}
	}
	return true
}
