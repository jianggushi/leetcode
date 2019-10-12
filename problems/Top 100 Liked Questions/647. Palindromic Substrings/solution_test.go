package solution

import "testing"

func Test_countSubstrings_1(t *testing.T) {
	s := "abc"
	expected := 3
	result := countSubstrings(s)
	if result != expected {
		t.Errorf("expected: %v, countSubstrings: %v", expected, result)
	}
}

func Test_countSubstrings_2(t *testing.T) {
	s := "aaa"
	expected := 6
	result := countSubstrings(s)
	if result != expected {
		t.Errorf("expected: %v, countSubstrings: %v", expected, result)
	}
}

func Test_countSubstrings_3(t *testing.T) {
	s := ""
	expected := 0
	result := countSubstrings(s)
	if result != expected {
		t.Errorf("expected: %v, countSubstrings: %v", expected, result)
	}
}

func Test_countSubstrings_4(t *testing.T) {
	s := "aabac"
	expected := 7
	result := countSubstrings(s)
	if result != expected {
		t.Errorf("expected: %v, countSubstrings: %v", expected, result)
	}
}
