package solution

import "testing"

func Test_isValid(t *testing.T) {
	s := ""
	if !isValid(s) {
		t.Error(s)
	}

	s = "()"
	if !isValid(s) {
		t.Error(s)
	}

	s = "()[]{}"
	if !isValid(s) {
		t.Error(s)
	}

	s = "{[]}"
	if !isValid(s) {
		t.Error(s)
	}

	s = "("
	if isValid(s) {
		t.Error(s)
	}

	s = "(]"
	if isValid(s) {
		t.Error(s)
	}

	s = "([)]"
	if isValid(s) {
		t.Error(s)
	}
}
