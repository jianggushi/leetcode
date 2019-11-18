package solution

import (
	"testing"
)

func Test_combinationSum4_1(t *testing.T) {
	nums := []int{1, 2, 3}
	target := 4
	expected := 7
	ans := combinationSum4(nums, target)
	if ans != expected {
		t.Errorf("expected: %v, combinationSum4: %v", expected, ans)
	}

}

func Test_combinationSum4_2(t *testing.T) {
	nums := []int{2, 1, 3}
	target := 35
	expected := 1132436852
	ans := combinationSum4(nums, target)
	if ans != expected {
		t.Errorf("expected: %v, combinationSum4: %v", expected, ans)
	}

}
