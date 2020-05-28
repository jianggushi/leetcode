package solution

import "testing"

func Test_decodeString_01(t *testing.T) {
	s := "3[a]2[bc]"
	expected := "aaabcbc"
	ans := decodeString(s)
	if ans != expected {
		t.Errorf("expected: %v, ans: %v", expected, ans)
	}
}

func Test_decodeString_02(t *testing.T) {
	s := "3[a2[c]]"
	expected := "accaccacc"
	ans := decodeString(s)
	if ans != expected {
		t.Errorf("expected: %v, ans: %v", expected, ans)
	}
}

func Test_decodeString_03(t *testing.T) {
	s := "2[abc]3[cd]ef"
	expected := "abcabccdcdcdef"
	ans := decodeString(s)
	if ans != expected {
		t.Errorf("expected: %v, ans: %v", expected, ans)
	}
}

func Test_decodeString_04(t *testing.T) {
	s := "2[a]"
	expected := "aa"
	ans := decodeString(s)
	if ans != expected {
		t.Errorf("expected: %v, ans: %v", expected, ans)
	}
}

func Test_decodeString_05(t *testing.T) {
	s := "a"
	expected := "a"
	ans := decodeString(s)
	if ans != expected {
		t.Errorf("expected: %v, ans: %v", expected, ans)
	}
}

func Test_decodeString_06(t *testing.T) {
	s := "[a]"
	expected := "a"
	ans := decodeString(s)
	if ans != expected {
		t.Errorf("expected: %v, ans: %v", expected, ans)
	}
}
