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

func Test_combine_1(t *testing.T) {
	n := 4
	k := 2
	expected := [][]int{
		[]int{2, 4},
		[]int{3, 4},
		[]int{2, 3},
		[]int{1, 2},
		[]int{1, 3},
		[]int{1, 4},
	}
	ans := combine(n, k)
	if !Equal(ans, expected) {
		t.Errorf("expected: %v, combine: %v", expected, ans)
	}
}
