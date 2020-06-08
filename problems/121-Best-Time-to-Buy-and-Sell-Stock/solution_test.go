package solution

import "testing"

func Test_maxProfit_01(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	expected := 5
	ans := maxProfit(prices)
	if ans != expected {
		t.Errorf("expected: %v, ans %v", expected, ans)
	}
}

func Test_maxProfit_02(t *testing.T) {
	prices := []int{7, 6, 4, 3, 1}
	expected := 0
	ans := maxProfit(prices)
	if ans != expected {
		t.Errorf("expected: %v, ans %v", expected, ans)
	}
}

func Test_maxProfit_03(t *testing.T) {
	prices := []int{}
	expected := 0
	ans := maxProfit(prices)
	if ans != expected {
		t.Errorf("expected: %v, ans %v", expected, ans)
	}
}

func Test_maxProfit_04(t *testing.T) {
	prices := []int{7}
	expected := 0
	ans := maxProfit(prices)
	if ans != expected {
		t.Errorf("expected: %v, ans %v", expected, ans)
	}
}

func Test_maxProfit_05(t *testing.T) {
	prices := []int{7}
	expected := 0
	ans := maxProfit(prices)
	if ans != expected {
		t.Errorf("expected: %v, ans %v", expected, ans)
	}
}

func Test_maxProfit_06(t *testing.T) {
	prices := []int{7, 2, 5, 1, 2}
	expected := 3
	ans := maxProfit(prices)
	if ans != expected {
		t.Errorf("expected: %v, ans %v", expected, ans)
	}
}

func Test_maxProfit_07(t *testing.T) {
	prices := []int{7, 2, 5, 1, 2, 5}
	expected := 4
	ans := maxProfit(prices)
	if ans != expected {
		t.Errorf("expected: %v, ans %v", expected, ans)
	}
}
