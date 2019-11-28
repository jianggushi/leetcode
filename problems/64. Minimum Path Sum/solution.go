package solution

// 解法二
// 动态规划
func minPathSum(grid [][]int) int {
	row := len(grid)
	col := len(grid[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i == 0 {
				grid[i][j] += grid[i][j-1]
				continue
			}
			if j == 0 {
				grid[i][j] += grid[i-1][j]
				continue
			}
			grid[i][j] += min(grid[i][j-1], grid[i-1][j])
		}
	}
	return grid[row-1][col-1]
}

// 解法一
// 回溯算法 + 记忆化
func minPathSum1(grid [][]int) int {
	memo := make(map[[2]int]int)
	return backtrack(grid, len(grid)-1, len(grid[0])-1, memo)
}

func backtrack(grid [][]int, row int, col int, memo map[[2]int]int) int {
	if row == 0 {
		t := 0
		for j := 0; j <= col; j++ {
			t += grid[0][j]
		}
		return t
	}
	if col == 0 {
		t := 0
		for i := 0; i <= row; i++ {
			t += grid[i][0]
		}
		return t
	}
	r1 := grid[row][col]
	if _, ok := memo[[2]int{row - 1, col}]; ok {
		r1 += memo[[2]int{row - 1, col}]
	} else {
		r1 += backtrack(grid, row-1, col, memo)
	}
	r2 := grid[row][col]
	if _, ok := memo[[2]int{row, col - 1}]; ok {
		r2 += memo[[2]int{row, col - 1}]
	} else {
		r2 += backtrack(grid, row, col-1, memo)
	}
	return min(r1, r2)
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}
