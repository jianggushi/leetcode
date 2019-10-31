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

func Test_combinationSum_1(t *testing.T) {
	candidates := []int{2, 3, 6, 7}
	target := 7
	expected := [][]int{
		[]int{2, 2, 3},
		[]int{7},
	}
	ans := combinationSum(candidates, target)
	if !Equal(ans, expected) {
		t.Errorf("expected: %v, combinationSum: %v", expected, ans)
	}
}

func Test_combinationSum_2(t *testing.T) {
	candidates := []int{2, 3, 5}
	target := 8
	expected := [][]int{
		[]int{2, 2, 2, 2},
		[]int{2, 3, 3},
		[]int{3, 5},
	}
	ans := combinationSum(candidates, target)
	if !Equal(ans, expected) {
		t.Errorf("expected: %v, combinationSum: %v", expected, ans)
	}
}
