package solution

import "testing"

func Test_findRepeatNumber_01(t *testing.T) {
	nums := []int{2, 3, 1, 0, 2, 5, 3}
	expected := map[int]struct{}{
		2: {},
		3: {},
	}
	ans := findRepeatNumber(nums)
	if _, ok := expected[ans]; !ok {
		t.Errorf("expected: %v, ans: %v", expected, ans)
	}
}
