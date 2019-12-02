package solution

import "testing"

func Test_numTrees_1(t *testing.T) {
	n := 3
	expected := 5
	ans := numTrees(n)
	if ans != expected {
		t.Errorf("expected: %v, numTrees: %v", expected, ans)
	}
}
