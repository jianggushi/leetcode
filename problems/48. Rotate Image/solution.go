package solution

func rotate(matrix [][]int) {
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

func rotate1(matrix [][]int) {
	n := len(matrix)
	ans := make([][]int, n)

	for i := 0; i < n; i++ {
		for j := n - 1; j >= 0; j-- {
			ans[i] = append(ans[i], matrix[j][i])
		}
	}
	matrix = ans
}
