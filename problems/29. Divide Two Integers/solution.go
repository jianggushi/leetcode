package solution

import "math"

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	if divisor == 0 {
		return math.MaxInt32
	}
	sign := 1
	if dividend > 0 && divisor < 0 {
		divisor = -divisor
		sign = -1
	}
	if dividend < 0 && divisor > 0 {
		dividend = -dividend
		sign = -1
	}
	ans := 1
	d := divisor
	for divisor <= dividend {
		divisor += d
		dividend -= d
		ans += 2
	}

	if divisor > dividend {
		ans--
	}
	if divisor-d > dividend {
		ans--
	}

	if sign == -1 {
		return -ans
	}

	return ans
}
