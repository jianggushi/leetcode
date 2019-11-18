package solution

import (
	"reflect"
	"testing"
)

func Equal(a [][]int, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	n := len(a)
	// a 中的元素在 b 中都可以找到
	for i := 0; i < n; i++ {
		find := false
		for j := 0; j < n; j++ {
			if reflect.DeepEqual(a[i], b[j]) {
				find = true
			}
		}
		if !find {
			return false
		}
	}
	// b 中的元素在 a 中都可以找到
	for i := 0; i < n; i++ {
		find := false
		for j := 0; j < n; j++ {
			if reflect.DeepEqual(b[i], a[j]) {
				find = true
			}
		}
		if !find {
			return false
		}
	}
	return true
}

func Test_subsets_1(t *testing.T) {
	nums := []int{1, 2, 3}
	expected := [][]int{
		{},
		{1},
		{2},
		{1, 2},
		{3},
		{1, 3},
		{2, 3},
		{1, 2, 3},
	}
	ans := subsets(nums)
	if !Equal(ans, expected) {
		t.Errorf("expected: %v, subsets: %v", expected, ans)
	}
}

func Test_subsets_2(t *testing.T) {
	nums := []int{9, 0, 3, 5, 7}
	expected := [][]int{
		{},
		{1},
		{2},
		{1, 2},
		{3},
		{1, 3},
		{2, 3},
		{1, 2, 3},
	}
	ans := subsets(nums)
	if !Equal(ans, expected) {
		t.Errorf("expected: %v, subsets: %v", expected, ans)
	}
}
