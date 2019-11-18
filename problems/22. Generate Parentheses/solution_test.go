package solution

import (
	"reflect"
	"testing"
)

func Equal(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	n := len(a)
	// a 中的元素在 b 中都可以找到
	for i := 0; i < n; i++ {
		find := false
		for j := 0; j < n; j++ {
			if reflect.DeepEqual(a[i], b[j]) {
				find = true
			}
		}
		if !find {
			return false
		}
	}
	// b 中的元素在 a 中都可以找到
	for i := 0; i < n; i++ {
		find := false
		for j := 0; j < n; j++ {
			if reflect.DeepEqual(b[i], a[j]) {
				find = true
			}
		}
		if !find {
			return false
		}
	}
	return true
}

func Test_generateParenthesis_1(t *testing.T) {
	n := 1
	expected := []string{
		"()",
	}
	ans := generateParenthesis(n)
	if !Equal(ans, expected) {
		t.Errorf("expected: %v, generateParenthesis: %v", expected, ans)
	}
}

func Test_generateParenthesis_2(t *testing.T) {
	n := 2
	expected := []string{
		"()()",
		"(())",
	}
	ans := generateParenthesis(n)
	if !Equal(ans, expected) {
		t.Errorf("expected: %v, generateParenthesis: %v", expected, ans)
	}
}

func Test_generateParenthesis_3(t *testing.T) {
	n := 3
	expected := []string{
		"((()))",
		"(()())",
		"(())()",
		"()(())",
		"()()()",
	}
	ans := generateParenthesis(n)
	if !Equal(ans, expected) {
		t.Errorf("expected: %v, generateParenthesis: %v", expected, ans)
	}
}

func Test_generateParenthesis_4(t *testing.T) {
	n := 0
	expected := []string{}
	ans := generateParenthesis(n)
	if !Equal(ans, expected) {
		t.Errorf("expected: %v, generateParenthesis: %v", expected, ans)
	}
}
