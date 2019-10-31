package solution

// 动态规划
// x & (x - 1) 使得 x 的最低有效位变为 0，那么
// f(x) = 1 + f(x & (x - 1))
// f(0) = 0
// f(1) = 1
func countBits(num int) []int {
	ans := make([]int, num+1)
	for i := 1; i <= num; i++ {
		ans[i] = 1 + ans[i&(i-1)]
	}
	return ans
}

// 两层循环 + 位操作小技巧
func countBits2(num int) []int {
	ans := make([]int, num+1)
	for i := 0; i <= num; i++ {
		bits := 0
		// x & (x - 1) 使得 x 的最低有效位变为 0
		for x := i; x != 0; {
			bits++
			x &= (x - 1)
		}
		ans[i] = bits
	}
	return ans
}

// 两层循环 + 位移
func countBits1(num int) []int {
	ans := make([]int, num+1)
	for i := 0; i <= num; i++ {
		bits := 0
		// 不断的右移，并 & 上 0x01
		for x := i; x != 0; x >>= 1 {
			if x&0x01 == 1 {
				bits++
			}
		}
		ans[i] = bits
	}
	return ans
}
