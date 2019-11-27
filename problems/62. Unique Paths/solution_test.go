package solution

import "testing"

func Test_uniquePaths_1(t *testing.T) {
	m, n := 3, 2
	expected := 3
	ans := uniquePaths(m, n)
	if ans != expected {
		t.Errorf("expected: %v, uniquePaths: %v", expected, ans)
	}
}

func Test_uniquePaths_2(t *testing.T) {
	m, n := 7, 3
	expected := 28
	ans := uniquePaths(m, n)
	if ans != expected {
		t.Errorf("expected: %v, uniquePaths: %v", expected, ans)
	}
}

func Test_uniquePaths_3(t *testing.T) {
	m, n := 1, 1
	expected := 1
	ans := uniquePaths(m, n)
	if ans != expected {
		t.Errorf("expected: %v, uniquePaths: %v", expected, ans)
	}
}

func Test_uniquePaths_4(t *testing.T) {
	m, n := 23, 12
	expected := 193536720
	ans := uniquePaths(m, n)
	if ans != expected {
		t.Errorf("expected: %v, uniquePaths: %v", expected, ans)
	}
}

func Test_uniquePaths_5(t *testing.T) {
	m, n := 51, 9
	expected := 1916797311
	ans := uniquePaths(m, n)
	if ans != expected {
		t.Errorf("expected: %v, uniquePaths: %v", expected, ans)
	}
}
