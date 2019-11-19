package solution

import "testing"

func Test_findDuplicate1(t *testing.T) {
	nums := []int{1, 3, 4, 2, 2}
	expected := 2
	ans := findDuplicate(nums)
	if ans != expected {
		t.Errorf("expected: %v, findDuplicate: %v", expected, ans)
	}
}

func Test_findDuplicate2(t *testing.T) {
	nums := []int{3, 1, 3, 4, 2}
	expected := 3
	ans := findDuplicate(nums)
	if ans != expected {
		t.Errorf("expected: %v, findDuplicate: %v", expected, ans)
	}
}
