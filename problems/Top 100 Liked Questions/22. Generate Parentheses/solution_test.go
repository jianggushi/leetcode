package solution

import "testing"

func Test_generateParenthesis_1(t *testing.T) {
	n := 1
	expected := 1
	ans := generateParenthesis(n)
	if len(ans) != expected {
		t.Errorf("expected: %v, generateParenthesis: %v", expected, ans)
		t.Errorf("generateParenthesis: %v", ans)
	}
}

func Test_generateParenthesis_2(t *testing.T) {
	n := 2
	expected := 2
	ans := generateParenthesis(n)
	if len(ans) != expected {
		t.Errorf("expected: %v, generateParenthesis: %v", expected, ans)
		t.Errorf("generateParenthesis: %v", ans)
	}
}

func Test_generateParenthesis_3(t *testing.T) {
	n := 3
	expected := 5
	ans := generateParenthesis(n)
	if len(ans) != expected {
		t.Errorf("expected: %v, generateParenthesis: %v", expected, ans)
		t.Errorf("generateParenthesis: %v", ans)
	}
}
