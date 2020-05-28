package solution

import (
	"strconv"
	"strings"
)

// 类似括号匹配，利用栈来解决
// s = "3[a]2[bc]", return "aaabcbc".
// s = "3[a2[c]]", return "accaccacc".
// s = "2[abc]3[cd]ef", return "abcabccdcdcdef".
func decodeString(s string) string {
	stack := make([]string, 0, len(s))
	n := len(s)
	for i := 0; i < n; {
		// 右括号
		if s[i] == ']' {
			i++
			ss := ""
			for len(stack) != 0 && stack[len(stack)-1] != "[" {
				ss = stack[len(stack)-1] + ss
				stack = stack[:len(stack)-1]
			}
			if ss != "" {
				stack = stack[:len(stack)-1] // 弹出栈顶的 "["
				if len(stack) != 0 {
					sd := stack[len(stack)-1] // 数字 k
					dd, _ := strconv.Atoi(sd)
					tmp := ss
					for i := 1; i < dd; i++ { // 重复 k 次 ss
						tmp += ss
					}
					ss = tmp
					stack[len(stack)-1] = ss // push 进去
				} else {
					return ss // 栈空了
				}
			}
		}
		if i >= n {
			break
		}
		// 左括号
		if s[i] == '[' {
			stack = append(stack, "[")
			i++
			continue
		}
		t := []byte{}
		// 数字
		for s[i] >= '0' && s[i] <= '9' {
			t = append(t, s[i])
			i++
		}
		if len(t) != 0 {
			stack = append(stack, string(t))
			continue
		}
		// 字符
		for i < n && s[i] != '[' && s[i] != ']' && !(s[i] >= '0' && s[i] <= '9') {
			t = append(t, s[i])
			i++
		}
		if len(t) != 0 {
			stack = append(stack, string(t))
			continue
		}
	}
	return strings.Join(stack, "")
}
