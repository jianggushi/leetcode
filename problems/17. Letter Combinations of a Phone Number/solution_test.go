package solution

import (
	"reflect"
	"testing"
)

func Test_letterCombinations_1(t *testing.T) {
	digits := "23"
	expected := []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}
	ans := letterCombinations(digits)
	if !reflect.DeepEqual(ans, expected) {
		t.Errorf("expected: %v, letterCombinations: %v", expected, ans)
	}
}

func Test_letterCombinations_2(t *testing.T) {
	digits := ""
	expected := []string{}
	ans := letterCombinations(digits)
	if !reflect.DeepEqual(ans, expected) {
		t.Errorf("expected: %v, letterCombinations: %v", expected, ans)
	}
}
