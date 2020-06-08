package solution

// 进一步简化 n == 1 时的特殊处理情况
func climbStairs(n int) int {
	dp1, dp2 := 0, 1
	for i := 1; i <= n; i++ {
		dp2, dp1 = dp2+dp1, dp2
	}
	return dp2
}

// 递归方式会重复计算
// 使用迭代方式
func climbStairs3(n int) int {
	if n == 1 {
		return 1
	}
	dp1, dp2 := 1, 1
	for i := 2; i <= n; i++ {
		dp := dp2 + dp1
		dp1 = dp2
		dp2 = dp
	}
	return dp2
}

// 纯递归解法，会重复计算
// 计算 f(6) = f(5) + f(4)
// 计算 f(5) = f(4) + f(3)
// 上面 f(4) 就产生了重复计算
// 记录计算过的，减少重复计算，提升效率
func climbStairs2(n int) int {
	mmap := make(map[int]int)
	return help(n, mmap)
}

func help(n int, mmap map[int]int) int {
	if n == 0 || n == 1 {
		return 1
	}
	if v, ok := mmap[n]; ok {
		return v
	}
	f1 := help(n-1, mmap)
	mmap[n-1] = f1
	f2 := help(n-2, mmap)
	mmap[n-2] = f2
	return f1 + f2
}

// 递归解法
// f(i) = f(i-1) + f(i-2)
// f(1) = 1
// f(0) = 1
func climbStairs1(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return climbStairs1(n-1) + climbStairs1(n-2)
}
