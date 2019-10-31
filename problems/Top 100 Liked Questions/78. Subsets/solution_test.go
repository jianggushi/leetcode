package solution

import (
	"reflect"
	"testing"
)

func Test_subsets_1(t *testing.T) {
	nums := []int{1, 2, 3}
	expected := [][]int{
		[]int{},
		[]int{1},
		[]int{2},
		[]int{1, 2},
		[]int{3},
		[]int{1, 3},
		[]int{2, 3},
		[]int{1, 2, 3},
	}
	ans := subsets(nums)
	if !reflect.DeepEqual(ans, expected) {
		t.Errorf("expected: %v, subsets: %v", expected, ans)
	}
}

func Test_subsets_2(t *testing.T) {
	nums := []int{9, 0, 3, 5, 7}
	expected := [][]int{
		[]int{},
		[]int{1},
		[]int{2},
		[]int{1, 2},
		[]int{3},
		[]int{1, 3},
		[]int{2, 3},
		[]int{1, 2, 3},
	}
	ans := subsets(nums)
	if !reflect.DeepEqual(ans, expected) {
		t.Errorf("expected: %v, subsets: %v", expected, ans)
	}
}
