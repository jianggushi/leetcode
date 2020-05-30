package solution

import "testing"

func Test_longestCommonPrefix_1(t *testing.T) {
	s := []string{"flower", "flow", "flight"}
	prefix := longestCommonPrefix(s)
	if prefix != "fl" {
		t.Error(prefix)
	}
}

func Test_longestCommonPrefix_2(t *testing.T) {
	s := []string{"dog", "racecar", "car"}
	prefix := longestCommonPrefix(s)
	if prefix != "" {
		t.Error(prefix)
	}
}
