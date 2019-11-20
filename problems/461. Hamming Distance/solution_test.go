package solution

import "testing"

func Test_hammingDistance1(t *testing.T) {
	x, y := 1, 4
	expected := 2
	ans := hammingDistance(x, y)
	if ans != expected {
		t.Errorf("expected: %v, hammingDistance: %v", expected, ans)
	}
}
