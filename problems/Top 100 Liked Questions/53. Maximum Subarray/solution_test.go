package solution

import "testing"

func Test_maxSubArray_1(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	expected := 6
	result := maxSubArray(nums)
	if result != expected {
		t.Errorf("expected: %v, get result: %v", expected, result)
	}
}

func Test_maxSubArray_2(t *testing.T) {
	nums := []int{-1}
	expected := -1
	result := maxSubArray(nums)
	if result != expected {
		t.Errorf("expected: %v, get result: %v", expected, result)
	}
}

func Test_maxSubArray_3(t *testing.T) {
	nums := []int{1, 2}
	expected := 3
	result := maxSubArray(nums)
	if result != expected {
		t.Errorf("expected: %v, get result: %v", expected, result)
	}
}
