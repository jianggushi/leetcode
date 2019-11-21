package solution

import "testing"

func Test_maxArea_1(t *testing.T) {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	expected := 49
	ans := maxArea(height)
	if ans != expected {
		t.Errorf("expected: %v, maxArea: %v", expected, ans)
	}
}
