package solution

import "testing"

func Test_divide1(t *testing.T) {
	dividend, divisor := 10, 3
	expected := 3
	ans := divide(dividend, divisor)
	if ans != expected {
		t.Errorf("expected: %v, divide: %v", expected, ans)
	}
}

func Test_divide2(t *testing.T) {
	dividend, divisor := 7, -3
	expected := -2
	ans := divide(dividend, divisor)
	if ans != expected {
		t.Errorf("expected: %v, divide: %v", expected, ans)
	}
}

func Test_divide3(t *testing.T) {
	dividend, divisor := 3, 3
	expected := 1
	ans := divide(dividend, divisor)
	if ans != expected {
		t.Errorf("expected: %v, divide: %v", expected, ans)
	}
}

func Test_divide4(t *testing.T) {
	dividend, divisor := -3, 3
	expected := -1
	ans := divide(dividend, divisor)
	if ans != expected {
		t.Errorf("expected: %v, divide: %v", expected, ans)
	}
}

func Test_divide5(t *testing.T) {
	dividend, divisor := 2, 3
	expected := 0
	ans := divide(dividend, divisor)
	if ans != expected {
		t.Errorf("expected: %v, divide: %v", expected, ans)
	}
}
