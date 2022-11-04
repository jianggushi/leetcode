// Source : https://leetcode.cn/problems/is-subsequence/
// Author : jianggushi
// Date   : 2022-11-02

// 分析
// 题目中说明
// 子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串
// ace是abcde的一个子序列，而aec不是

package solution

// 双指针写法简化代码
// 初始两个指针i和j分别指向s和t的开始位置
// 匹配s[i]和t[j]
// 如果匹配成功，i和j同时向后移动，匹配s的下一个位置
// 如果匹配失败，j向后移动，i不变，尝试用t的下一个位置进行匹配
// 怎样算全部成功匹配了呢？
// i到达s的末尾且j没到t的末尾或同时到末尾
func isSubsequence(s string, t string) bool {
	m, n := len(s), len(t)
	i, j := 0, 0
	for i < m && j < n {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == m
}

// 解法一：遍历法
//
// 从左向右遍历字符串s
// 找到s[i]在t中出现的位置p，如果没有出现说明s不是t的子序列
// 如果找到了，下次从p+1开始寻找s[i+1]在t中出现的位置
// 直到遍历完字符串s
//
// 时间复杂度：O(m+n)
// 空间复杂度：O(1)
func isSubsequence1(s string, t string) bool {
	m, n := len(s), len(t)
	p := 0
	for i := 0; i < m; i++ {
		match := false
		for j := p; j < n; j++ {
			if s[i] == t[j] {
				match = true
				p = j + 1
				break
			}
		}
		if !match {
			return false
		}
	}
	return true
}
