package solution

import "testing"

func Test_longestPalindrome_1(t *testing.T) {
	s := "babad"
	expected := "bab"
	ans := longestPalindrome(s)
	if ans != expected {
		t.Errorf("expected: %v, longestPalindrome: %v", expected, ans)
	}
}
func Test_longestPalindrome_2(t *testing.T) {
	s := "cbbd"
	expected := "bb"
	ans := longestPalindrome(s)
	if ans != expected {
		t.Errorf("expected: %v, longestPalindrome: %v", expected, ans)
	}
}

func Test_longestPalindrome_3(t *testing.T) {
	s := ""
	expected := ""
	ans := longestPalindrome(s)
	if ans != expected {
		t.Errorf("expected: %v, longestPalindrome: %v", expected, ans)
	}
}

func Test_longestPalindrome_4(t *testing.T) {
	s := "a"
	expected := "a"
	ans := longestPalindrome(s)
	if ans != expected {
		t.Errorf("expected: %v, longestPalindrome: %v", expected, ans)
	}
}
