package solution

// Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
// An input string is valid if:
// 1. Open brackets must be closed by the same type of brackets.
// 2. Open brackets must be closed in the correct order.
// Note that an empty string is also considered valid.

func isValid1(s string) bool {
	var stack []byte

	l := len(s)
	if l == 0 {
		return true
	}

	for i := 0; i < l; i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 {
				return false
			}
			if (s[i] == ')' && stack[len(stack)-1] != '(') ||
				(s[i] == '}' && stack[len(stack)-1] != '{') ||
				(s[i] == ']' && stack[len(stack)-1] != '[') {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}

func isValid(s string) bool {
	var stack []byte
	mapping := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}

	l := len(s)
	if l == 0 {
		return true
	}

	for i := 0; i < l; i++ {
		// 右括号：查看 stack 中是否是对应的左括号
		if v, ok := mapping[s[i]]; ok {
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] != v {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			// 左括号：直接压栈
			stack = append(stack, s[i])
		}
	}
	// NOTE: 遍历完后，需要检查下 stack 是否为空
	if len(stack) != 0 {
		return false
	}
	return true
}
