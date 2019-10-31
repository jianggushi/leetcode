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

func Test_combinationSum3_1(t *testing.T) {
	k := 3
	n := 7
	expected := [][]int{
		[]int{1, 2, 4},
	}
	ans := combinationSum3(k, n)
	if !Equal(ans, expected) {
		t.Errorf("expected: %v, combinationSum3:%v", expected, ans)
	}
}

func Test_combinationSum3_2(t *testing.T) {
	k := 3
	n := 9
	expected := [][]int{
		[]int{1, 2, 6},
		[]int{1, 3, 5},
		[]int{2, 3, 4},
	}
	ans := combinationSum3(k, n)
	if !Equal(ans, expected) {
		t.Errorf("expected: %v, combinationSum3:%v", expected, ans)
	}
}
