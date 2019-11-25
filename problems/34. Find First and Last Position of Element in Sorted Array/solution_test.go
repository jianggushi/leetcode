package solution

import (
	"reflect"
	"testing"
)

func Test_searchRange_1(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	expected := []int{3, 4}
	ans := searchRange(nums, target)
	if !reflect.DeepEqual(ans, expected) {
		t.Errorf("expected: %v, searchRange: %v", expected, ans)
	}
}

func Test_searchRange_2(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 6
	expected := []int{-1, -1}
	ans := searchRange(nums, target)
	if !reflect.DeepEqual(ans, expected) {
		t.Errorf("expected: %v, searchRange: %v", expected, ans)
	}
}
