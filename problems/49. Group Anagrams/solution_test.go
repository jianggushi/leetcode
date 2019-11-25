package solution

import (
	"reflect"
	"testing"
)

func Test_groupAnagrams_1(t *testing.T) {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	expected := [][]string{
		{"ate", "eat", "tea"},
		{"nat", "tan"},
		{"bat"},
	}
	ans := groupAnagrams(strs)
	if !reflect.DeepEqual(ans, expected) {
		t.Errorf("expected: %v, groupAnagrams: %v", expected, ans)
	}
}
