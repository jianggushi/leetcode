package solution

// 解法六
// 动态规划 + 空间压缩（一维动态规划）
// 在解法五的基础上进一步压缩空间，只使用一个数组完成复用
func uniquePaths(m int, n int) int {
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] += dp[j-1]
		}
	}
	return dp[n-1]
}

// 解法五
// 动态规划 + 空间压缩
// 当前位置的 dp 值只需要左边一个和上边一个就可以得到，没必要把所有的 dp 值都存下来
// 可以使用两个长为 n 的数组滚动复用
func uniquePaths5(m int, n int) int {
	dp1 := make([]int, n)
	dp2 := make([]int, n)
	for i := 0; i < n; i++ {
		dp1[i] = 1
		dp2[i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp2[j] = dp2[j-1] + dp1[j]
		}
		// 滚动复用两个长为 n 的数组
		dp1, dp2 = dp2, dp1
	}
	return dp1[n-1]
}

// 解法四
// 动态规划（二维动态规划）
func uniquePaths4(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = 1
		if i == 0 {
			for j := 0; j < n; j++ {
				dp[i][j] = 1
			}
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

// 解法三
// 回溯算法 + 记忆化
func uniquePaths3(m int, n int) int {
	memo := make(map[[2]int]int)
	return path(m, n, memo)
}

func path(m int, n int, memo map[[2]int]int) int {
	if m == 1 || n == 1 {
		return 1
	}
	r1, r2 := 0, 0
	if val, ok := memo[[2]int{m - 1, n}]; ok {
		r1 = val
	} else {
		r1 = path(m-1, n, memo)
		memo[[2]int{m - 1, n}] = r1
	}
	if val, ok := memo[[2]int{m, n - 1}]; ok {
		r2 = val
	} else {
		r2 = path(m, n-1, memo)
		memo[[2]int{m, n - 1}] = r2
	}
	return r1 + r2
}

// 解法二
// 回溯算法，简化代码
func uniquePaths2(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}
	return uniquePaths2(m-1, n) + uniquePaths2(m, n-1)
}

// 解法一
// 回溯算法
func uniquePaths1(m int, n int) int {
	ans := 0
	backtrack(m, n, &ans)
	return ans
}

func backtrack(m int, n int, ans *int) {
	if m == 1 && n == 1 {
		(*ans)++
		return
	}
	if m > 0 {
		backtrack(m-1, n, ans)
	}
	if n > 0 {
		backtrack(m, n-1, ans)
	}
}
