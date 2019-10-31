package solution

import (
	"reflect"
	"testing"
)

func Test_twoSum_1(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	expected := []int{0, 1}
	result := twoSum(nums, target)
	eq := reflect.DeepEqual(result, expected)
	if !eq {
		t.Errorf("expected: %v, toSum: %v", expected, result)
	}
}

func Test_twoSum_2(t *testing.T) {
	nums := []int{3, 2, 4}
	target := 6
	expected := []int{1, 2}
	result := twoSum(nums, target)
	eq := reflect.DeepEqual(result, expected)
	if !eq {
		t.Errorf("expected: %v, toSum: %v", expected, result)
	}
}
