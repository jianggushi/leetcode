package solution

import "testing"

func Test_findUnsortedSubarray_1(t *testing.T) {
	nums := []int{2, 6, 4, 8, 10, 9, 15}
	expected := 5

	result := findUnsortedSubarray(nums)
	if result != expected {
		t.Errorf("expected: %v, findUnsortedSubarray: %v", expected, result)
	}
}

func Test_findUnsortedSubarray_2(t *testing.T) {
	nums := []int{2, 2, 2, 2}
	expected := 0

	result := findUnsortedSubarray(nums)
	if result != expected {
		t.Errorf("expected: %v, findUnsortedSubarray: %v", expected, result)
	}
}

func Test_findUnsortedSubarray_3(t *testing.T) {
	nums := []int{2}
	expected := 0

	result := findUnsortedSubarray(nums)
	if result != expected {
		t.Errorf("expected: %v, findUnsortedSubarray: %v", expected, result)
	}
}

func Test_findUnsortedSubarray_4(t *testing.T) {
	nums := []int{}
	expected := 0

	result := findUnsortedSubarray(nums)
	if result != expected {
		t.Errorf("expected: %v, findUnsortedSubarray: %v", expected, result)
	}
}

func Test_findUnsortedSubarray_5(t *testing.T) {
	expected := 0

	result := findUnsortedSubarray(nil)
	if result != expected {
		t.Errorf("expected: %v, findUnsortedSubarray: %v", expected, result)
	}
}
