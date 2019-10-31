package solution

import "testing"

func Test_rob_1(t *testing.T) {
	nums := []int{1, 2, 3, 1}
	expected := 4
	result := rob(nums)
	if result != expected {
		t.Errorf("expected: %v, rob: %v", expected, result)
	}
}

func Test_rob_2(t *testing.T) {
	nums := []int{2, 7, 9, 3, 1}
	expected := 12
	result := rob(nums)
	if result != expected {
		t.Errorf("expected: %v, rob: %v", expected, result)
	}
}

func Test_rob_3(t *testing.T) {
	nums := []int{}
	expected := 0
	result := rob(nums)
	if result != expected {
		t.Errorf("expected: %v, rob: %v", expected, result)
	}
}
