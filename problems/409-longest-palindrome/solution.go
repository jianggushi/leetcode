// Source : https://leetcode.cn/problems/longest-palindrome/
// Author : jianggushi
// Date   : 2022-11-15

package solution

// 解法三：进一步优化
//
// 字符可能出现奇数次或者偶数次
// 这和最终回文串的长度有很大关系
// 如果有字符出现奇数次，那么最终回文串的长度也为奇数
// 如果全部字符出现偶数次，那么最终回文串的长度也为偶数
// 利用这个特点简化代码
// 如果有字符出现奇数次，而且次数回文串的长度为偶数，那么回文串长度加1
func longestPalindrome(s string) int {
	m := ['z' - 'A' + 1]int{}
	for i := range s {
		m[s[i]-'A']++
	}
	ans := 0 // ans 回文串长度
	for _, v := range m {
		ans += (v >> 1) << 1
		// ans += (v / 2) * 2
		if v%2 == 1 && ans%2 == 0 { // 字符出现奇数次且当前回文串长度为偶数
			ans++
		}
	}
	return ans
}

// 解法二：数组
//
// 题目中限定s中只有大小写英文字符a-z和A-Z
// 这些字符都是编码在128以内的ASCII字符
// 可以用数组记录字符出现的次数
// 还可以利用A-Z和a-z的ASCII码连续性
// 进一步减少数组大小
func longestPalindrome2(s string) int {
	m := ['z' - 'A' + 1]int{}
	for i := range s {
		m[s[i]-'A']++
	}
	ans, l := 0, 0 // ans 回文串长度，l 有字符出现奇数次
	for _, v := range m {
		if v%2 == 0 {
			ans += v
		} else {
			// ans += (v / 2) * 2
			ans += (v - 1)
			l = 1
		}
	}
	return ans + l
}

// 解法一：哈希表
//
// 使用哈希表记录每个字符出现的次数
// 如果一个字符出现偶数次，那么这个字符全部可以用来构建回文串
// 如果一个字符出现奇数次，那么要取其中的偶数次用来构建回文串
// 最后如果字符中存在出现奇数次的情况，那么还能用一个字符构建回文串，回文串长度+1
//
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func longestPalindrome1(s string) int {
	m := map[byte]int{}
	for i := range s {
		if m[s[i]] > 0 {
			m[s[i]]++
		} else {
			m[s[i]] = 1
		}
	}
	ans, l := 0, 0 // ans 回文串长度，l 有字符出现奇数次
	for _, v := range m {
		if v%2 == 0 {
			ans += v
		} else {
			// ans += (v / 2) * 2
			ans += (v - 1)
			l = 1
		}
	}
	return ans + l
}
