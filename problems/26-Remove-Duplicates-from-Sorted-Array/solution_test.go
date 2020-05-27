package solution

import "testing"

func Test_removeDuplicates_01(t *testing.T) {
	nums := []int{1, 1, 2}
	expected := 2
	ans := removeDuplicates(nums)
	if ans != expected {
		t.Errorf("expected: %v, ans: %v", expected, ans)
	}
}

func Test_removeDuplicates_02(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	expected := 5
	ans := removeDuplicates(nums)
	if ans != expected {
		t.Errorf("expected: %v, ans: %v", expected, ans)
	}
}

func Test_removeDuplicates_03(t *testing.T) {
	nums := []int{0, 0, 0}
	expected := 1
	ans := removeDuplicates(nums)
	if ans != expected {
		t.Errorf("expected: %v, ans: %v", expected, ans)
	}
}

func Test_removeDuplicates_04(t *testing.T) {
	nums := []int{0, 1, 2}
	expected := 3
	ans := removeDuplicates(nums)
	if ans != expected {
		t.Errorf("expected: %v, ans: %v", expected, ans)
	}
}

func Test_removeDuplicates_05(t *testing.T) {
	nums := []int{0, 1, 1}
	expected := 2
	ans := removeDuplicates(nums)
	if ans != expected {
		t.Errorf("expected: %v, ans: %v", expected, ans)
	}
}
