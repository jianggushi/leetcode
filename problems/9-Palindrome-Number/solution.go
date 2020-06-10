package solution

import "strconv"

// 反转一半数字
// 特殊处理个位数为 0 的情况
func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	y := 0
	for x > y {
		y = y*10 + x%10
		x = x / 10
	}
	if x == y || x == y/10 {
		return true
	}
	return false
}

// 反转数字
// 比较反转后的数字和原始数字
// 溢出了怎么办？
func isPalindrome2(x int) bool {
	if x < 0 {
		return false
	}
	y := 0
	for i := x; i != 0; i /= 10 {
		y = y*10 + i%10
	}
	return x == y
}

// 把数字转成字符串
// 判断字符串是否为回文字符串
func isPalindrome1(x int) bool {
	s := strconv.Itoa(x)
	n := len(s)
	i, j := 0, n-1
	for i <= j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
