package solution

// 解法一：
// 异或操作可以把 x 和 y 中 bit 位不同的变成 1, bit 位相同的变成 0
// 问题就变为计算 x ^ y 后 bit 1 的个数

// 利用 n & (n-1) 会把 n 的最低位 1 变成 0 的特点，计算 bit 1 的个数
func hammingDistance(x int, y int) int {
	n := x ^ y
	bits := 0
	for n != 0 {
		bits++
		n &= (n - 1)
	}
	return bits
}

// 利用循环+移位的方式, 计算 bit 1 的个数
func hammingDistance1(x int, y int) int {
	n := x ^ y
	bits := 0
	for ; n != 0; n >>= 1 {
		if n&1 == 1 {
			bits++
		}
	}
	return bits
}
