package solution

import "testing"

func Test_search_1(t *testing.T) {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	expected := 4
	ans := search(nums, target)
	if ans != expected {
		t.Errorf("expected: %v, search: %v", expected, ans)
	}
}

func Test_search_2(t *testing.T) {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 4
	expected := 0
	ans := search(nums, target)
	if ans != expected {
		t.Errorf("expected: %v, search: %v", expected, ans)
	}
}

func Test_search_3(t *testing.T) {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 2
	expected := 6
	ans := search(nums, target)
	if ans != expected {
		t.Errorf("expected: %v, search: %v", expected, ans)
	}
}

func Test_search_4(t *testing.T) {
	nums := []int{}
	target := 5
	expected := -1
	ans := search(nums, target)
	if ans != expected {
		t.Errorf("expected: %v, search: %v", expected, ans)
	}
}

func Test_search_5(t *testing.T) {
	nums := []int{1}
	target := 0
	expected := -1
	ans := search(nums, target)
	if ans != expected {
		t.Errorf("expected: %v, search: %v", expected, ans)
	}
}

func Test_search_6(t *testing.T) {
	nums := []int{1, 3}
	target := 0
	expected := -1
	ans := search(nums, target)
	if ans != expected {
		t.Errorf("expected: %v, search: %v", expected, ans)
	}
}

func Test_search_7(t *testing.T) {
	nums := []int{1}
	target := 1
	expected := 0
	ans := search(nums, target)
	if ans != expected {
		t.Errorf("expected: %v, search: %v", expected, ans)
	}
}

func Test_search_8(t *testing.T) {
	nums := []int{3, 1}
	target := 1
	expected := 1
	ans := search(nums, target)
	if ans != expected {
		t.Errorf("expected: %v, search: %v", expected, ans)
	}
}

func Test_search_9(t *testing.T) {
	nums := []int{3, 4, 5, 1, 2}
	target := 4
	expected := 1
	ans := search(nums, target)
	if ans != expected {
		t.Errorf("expected: %v, search: %v", expected, ans)
	}
}

func Test_search_12(t *testing.T) {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 3
	expected := -1
	ans := search(nums, target)
	if ans != expected {
		t.Errorf("expected: %v, search: %v", expected, ans)
	}
}
