package solution

import (
	"math"
)

func myAtoi(str string) int {
	n := len(str)
	ans := int32(0)
	sign := int32(1)
	i := 0
	// 去除前面的空白
	for i < n && str[i] == ' ' {
		i++
	}
	// 正负号
	if i < n && (str[i] == '-' || str[i] == '+') {
		if str[i] == '-' {
			sign = -1
		}
		i++
	}
	for i < n {
		if str[i] < '0' || str[i] > '9' {
			break
		}
		if ans > 214748364 {
			if sign == 1 {
				ans = math.MaxInt32
			} else {
				ans = math.MinInt32
			}
			break
		}
		if ans == 214748364 {
			if sign == 1 && (str[i]-'0' > 7) {
				ans = math.MaxInt32
				break
			} else if sign == -1 && (str[i]-'0' > 8) {
				ans = math.MinInt32
				break
			}
		}
		ans = ans*10 + int32(str[i]-'0')
		i++
	}
	return int(ans * sign)
}
