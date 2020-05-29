package solution

// 动态规划
// 推导状态转移方程，先从最简单的情况开始：
// 记 dp(i) = 从前 i 个房屋中能抢劫到的最大数额，Ai = 第 i 个房屋的钱数。
// 当 i = 1 时，有多少抢多少，dp(1) = A1
// 当 i = 2 时，哪个多抢哪个，dp(2) = max(A1, A2)
// 当 i = 3 时，权衡下抢不抢，如果抢得到 dp(1) + A3，如果不抢得到 dp(2)，那么 dp(3) = max(dp(2), dp(1)+A3)
// 以此类推，当 i = n 时
// dp(n) = max(dp(n-1), dp(n-2)+An)

// 迭代方式，再简化代码
// n == 0 这种特殊情况的判断也可以去掉
// 循环中通过临时变量的赋值方式，也可以简化掉
func rob4(nums []int) int {
	n := len(nums)
	dp0, dp1 := 0, 0
	for i := 0; i < n; i++ {
		dp0, dp1 = dp1, max(dp1, dp0+nums[i])
	}
	return dp1
}

// 迭代方式，简化代码
// 从递推公式看，需要 dp(n-1) 和 dp(n-2) 才能得到 dp(n)
// 为了求解 dp(2) 和 dp(1) 需要 dp(0) 和 dp(-1)
// 可以假设 dp(0) =0，dp(-1) = 0
// 这样可以把求解 dp(2) 和 dp(1) 这两种特殊情况也归纳进去
func rob3(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp0 := 0
	dp1 := 0
	for i := 0; i < n; i++ {
		dp2 := max(dp1, dp0+nums[i])
		dp0 = dp1
		dp1 = dp2
	}
	return dp1
}

// 动态规划
func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	dp0 := 0
	dp1 := nums[0]
	for i := 1; i < n; i++ {
		dp2 := max(dp1, dp0+nums[i])
		dp0 = dp1
		dp1 = dp2
	}
	return dp1
}

// 递归方式
// 使用一个 map 辅助记忆之前计算过的值
// 递归前先从 map 中查找
func rob2(nums []int) int {
	mmap := make(map[int]int)
	return calc(nums, mmap)
}

func calc(nums []int, mmap map[int]int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	if v, ok := mmap[n]; ok {
		return v
	}
	res := max(calc(nums[:n-1], mmap), calc(nums[:n-2], mmap)+nums[n-1])
	mmap[n] = res
	return res
}

// 递归方式
// 要解决一个大问题，先考虑解决一个小问题
// rob(n) = max(rob(n-1), rob(n-2) + n)
// 时间复杂度有点高，因为存在重复计算
func rob1(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	return max(rob1(nums[:n-1]), rob1(nums[:n-2])+nums[n-1])
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
