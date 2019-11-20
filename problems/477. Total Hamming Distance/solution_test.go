package solution

import "testing"

func Test_totalHammingDistance1(t *testing.T) {
	nums := []int{4, 14, 2}
	expected := 6
	ans := totalHammingDistance(nums)
	if ans != expected {
		t.Errorf("expected: %v, totalHammingDistance: %v", expected, ans)
	}
}
