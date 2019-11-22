package solution

import (
	"reflect"
	"testing"
)

func Test_nextPermutation_1(t *testing.T) {
	nums := []int{1, 2, 3}
	expected := []int{1, 3, 2}
	nextPermutation(nums)
	if !reflect.DeepEqual(nums, expected) {
		t.Errorf("expected: %v, nextPermutation: %v", expected, nums)
	}
}

func Test_nextPermutation_2(t *testing.T) {
	nums := []int{1, 3, 2}
	expected := []int{2, 1, 3}
	nextPermutation(nums)
	if !reflect.DeepEqual(nums, expected) {
		t.Errorf("expected: %v, nextPermutation: %v", expected, nums)
	}
}

func Test_nextPermutation_3(t *testing.T) {
	nums := []int{2, 1, 3}
	expected := []int{2, 3, 1}
	nextPermutation(nums)
	if !reflect.DeepEqual(nums, expected) {
		t.Errorf("expected: %v, nextPermutation: %v", expected, nums)
	}
}

func Test_nextPermutation_4(t *testing.T) {
	nums := []int{2, 3, 1}
	expected := []int{3, 1, 2}
	nextPermutation(nums)
	if !reflect.DeepEqual(nums, expected) {
		t.Errorf("expected: %v, nextPermutation: %v", expected, nums)
	}
}

func Test_nextPermutation_5(t *testing.T) {
	nums := []int{3, 1, 2}
	expected := []int{3, 2, 1}
	nextPermutation(nums)
	if !reflect.DeepEqual(nums, expected) {
		t.Errorf("expected: %v, nextPermutation: %v", expected, nums)
	}
}

func Test_nextPermutation_6(t *testing.T) {
	nums := []int{3, 2, 1}
	expected := []int{1, 2, 3}
	nextPermutation(nums)
	if !reflect.DeepEqual(nums, expected) {
		t.Errorf("expected: %v, nextPermutation: %v", expected, nums)
	}
}

func Test_nextPermutation_7(t *testing.T) {
	nums := []int{1, 1, 5}
	expected := []int{1, 5, 1}
	nextPermutation(nums)
	if !reflect.DeepEqual(nums, expected) {
		t.Errorf("expected: %v, nextPermutation: %v", expected, nums)
	}
}

func Test_nextPermutation_8(t *testing.T) {
	nums := []int{1}
	expected := []int{1}
	nextPermutation(nums)
	if !reflect.DeepEqual(nums, expected) {
		t.Errorf("expected: %v, nextPermutation: %v", expected, nums)
	}
}

func Test_nextPermutation_9(t *testing.T) {
	nums := []int{1, 2}
	expected := []int{2, 1}
	nextPermutation(nums)
	if !reflect.DeepEqual(nums, expected) {
		t.Errorf("expected: %v, nextPermutation: %v", expected, nums)
	}
}

func Test_nextPermutation_10(t *testing.T) {
	nums := []int{1, 1, 3, 0}
	expected := []int{1, 3, 0, 1}
	nextPermutation(nums)
	if !reflect.DeepEqual(nums, expected) {
		t.Errorf("expected: %v, nextPermutation: %v", expected, nums)
	}
}

func Test_nextPermutation_11(t *testing.T) {
	nums := []int{5, 4, 7, 5, 3, 2}
	expected := []int{5, 5, 2, 3, 4, 7}
	nextPermutation(nums)
	if !reflect.DeepEqual(nums, expected) {
		t.Errorf("expected: %v, nextPermutation: %v", expected, nums)
	}
}
