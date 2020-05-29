package solution

import (
	"math"
)

// 注意题目中限制了是 int32 [-2^31, 2^31-1]
// 如果结果溢出，返回 0
// 如果直接操作数字的话，有可能溢出，怎么判断是否有溢出
func reverse(x int) int {
	var ans int
	max := math.MaxInt32 / 10
	min := math.MinInt32 / 10
	for x != 0 { // 迭代反转
		if ans > max || (ans == max && x%10 > 7) {
			return 0
		}
		if ans < min || (ans == min && x%10 < -8) {
			return 0
		}
		ans = ans*10 + x%10
		x = x / 10
	}
	return ans
}
