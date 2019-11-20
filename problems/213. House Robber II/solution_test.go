package solution

import "testing"

func Test_rob_1(t *testing.T) {
	nums := []int{2, 3, 2}
	expected := 3
	ans := rob(nums)
	if ans != expected {
		t.Errorf("expected: %v, rob: %v", expected, ans)
	}
}

func Test_rob_2(t *testing.T) {
	nums := []int{1, 2, 3, 1}
	expected := 4
	ans := rob(nums)
	if ans != expected {
		t.Errorf("expected: %v, rob: %v", expected, ans)
	}
}
