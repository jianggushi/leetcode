package solution

import (
	"reflect"
	"testing"
)

func Test_sortColors_1(t *testing.T) {
	nums := []int{2, 0, 2, 1, 1, 0}
	expected := []int{0, 0, 1, 1, 2, 2}
	sortColors(nums)
	if !reflect.DeepEqual(nums, expected) {
		t.Errorf("expected: %v, sortColors: %v", expected, nums)
	}
}

func Test_sortColors_2(t *testing.T) {
	nums := []int{0, 0, 2, 1, 1, 2}
	expected := []int{0, 0, 1, 1, 2, 2}
	sortColors(nums)
	if !reflect.DeepEqual(nums, expected) {
		t.Errorf("expected: %v, sortColors: %v", expected, nums)
	}
}

func Test_sortColors_3(t *testing.T) {
	nums := []int{2, 2, 0, 0, 1, 1}
	expected := []int{0, 0, 1, 1, 2, 2}
	sortColors(nums)
	if !reflect.DeepEqual(nums, expected) {
		t.Errorf("expected: %v, sortColors: %v", expected, nums)
	}
}

func Test_sortColors_4(t *testing.T) {
	nums := []int{0, 0}
	expected := []int{0, 0}
	sortColors(nums)
	if !reflect.DeepEqual(nums, expected) {
		t.Errorf("expected: %v, sortColors: %v", expected, nums)
	}
}

func Test_sortColors_5(t *testing.T) {
	nums := []int{1, 1}
	expected := []int{1, 1}
	sortColors(nums)
	if !reflect.DeepEqual(nums, expected) {
		t.Errorf("expected: %v, sortColors: %v", expected, nums)
	}
}
