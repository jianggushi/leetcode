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

func Test_combinationSum2_1(t *testing.T) {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	expected := [][]int{
		[]int{1, 7},
		[]int{1, 2, 5},
		[]int{2, 6},
		[]int{1, 1, 6},
	}
	ans := combinationSum2(candidates, target)
	if !Equal(ans, expected) {
		t.Errorf("expected: %v, combinationSum2: %v", expected, ans)
	}
}
