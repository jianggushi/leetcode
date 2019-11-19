package solution

import "testing"

func Test_isHappy1(t *testing.T) {
	expected := true
	ans := isHappy(19)
	if ans != expected {
		t.Errorf("expected: %v, isHappy: %v", expected, ans)
	}
}
