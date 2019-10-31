package solution

import (
	"reflect"
	"testing"
)

func Test_combinationSum_1(t *testing.T) {
	candidates := []int{2, 3, 6, 7}
	target := 7
	expected := [][]int{
		[]int{2, 2, 3},
		[]int{7},
	}
	ans := combinationSum(candidates, target)
	if !reflect.DeepEqual(ans, expected) {
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
	if !reflect.DeepEqual(ans, expected) {
		t.Errorf("expected: %v, combinationSum: %v", expected, ans)
	}
}
