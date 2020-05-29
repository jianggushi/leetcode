package solution

func rotate2(matrix [][]int) {
	n := len(matrix)

	for col := 0; col < n; col++ {
		for row := n - 1; row >= min(col+1, n-1-col); row-- {
			matrix[row][col], matrix[col][n-1-row] = matrix[col][n-1-row], matrix[row][col]
		}
	}
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

// 额外数组
func rotate(matrix [][]int) {
	n := len(matrix)
	ans := make([][]int, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]int, n)
		for j := 0; j < n; j++ {
			ans[i][j] = matrix[n-1-j][i]
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			matrix[i][j] = ans[i][j]
		}
	}
}
