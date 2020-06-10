package solution

import "testing"

func Test_subarraysDivByK_01(t *testing.T) {
	A := []int{4, 5, 0, -2, -3, 1}
	K := 5
	expected := 7
	ans := subarraysDivByK(A, K)
	if ans != expected {
		t.Errorf("expected: %v, ans: %v", expected, ans)
	}
}

func Test_subarraysDivByK_02(t *testing.T) {
	A := []int{4}
	K := 5
	expected := 0
	ans := subarraysDivByK(A, K)
	if ans != expected {
		t.Errorf("expected: %v, ans: %v", expected, ans)
	}
}

func Test_subarraysDivByK_03(t *testing.T) {
	A := []int{4, 5}
	K := 5
	expected := 1
	ans := subarraysDivByK(A, K)
	if ans != expected {
		t.Errorf("expected: %v, ans: %v", expected, ans)
	}
}

func Test_subarraysDivByK_04(t *testing.T) {
	A := []int{5, 5}
	K := 5
	expected := 3
	ans := subarraysDivByK(A, K)
	if ans != expected {
		t.Errorf("expected: %v, ans: %v", expected, ans)
	}
}

func Test_subarraysDivByK_05(t *testing.T) {
	A := []int{-1, 2, 9}
	K := 2
	expected := 2
	ans := subarraysDivByK(A, K)
	if ans != expected {
		t.Errorf("expected: %v, ans: %v", expected, ans)
	}
}

func Test_subarraysDivByK_06(t *testing.T) {
	A := []int{-6, 1, -5, 10}
	K := 9
	expected := 1
	ans := subarraysDivByK(A, K)
	if ans != expected {
		t.Errorf("expected: %v, ans: %v", expected, ans)
	}
}
