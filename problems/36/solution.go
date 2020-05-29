package solution

func isValidSudoku(board [][]byte) bool {
	n := len(board)
	ans := true
	for i := 0; i < n; i++ {
		ans = ans && isValidRaw(board, i)
		if !ans {
			return ans
		}
	}
	for i := 0; i < n; i++ {
		ans = ans && isValidCol(board, i)
		if !ans {
			return ans
		}
	}
	for i := 0; i < n; i = i + 3 {
		for j := 0; j < n; j = j + 3 {
			ans = ans && isValidBox(board, i, j)
			if !ans {
				return ans
			}
		}
	}
	return ans
}

func isValidRaw(board [][]byte, row int) bool {
	n := len(board[row])
	mmap := make(map[byte]int)
	for i := 0; i < n; i++ {
		if board[row][i] != '.' {
			if mmap[board[row][i]] > 0 {
				return false
			}
			mmap[board[row][i]]++
		}
	}
	return true
}

func isValidCol(board [][]byte, col int) bool {
	n := len(board)
	mmap := make(map[byte]int)
	for i := 0; i < n; i++ {
		if board[i][col] != '.' {
			if mmap[board[i][col]] > 0 {
				return false
			}
			mmap[board[i][col]]++
		}
	}
	return true
}

func isValidBox(board [][]byte, row, col int) bool {
	mmap := make(map[byte]int)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[row+i][col+j] != '.' {
				if mmap[board[row+i][col+j]] > 0 {
					return false
				}
				mmap[board[row+i][col+j]]++
			}
		}
	}
	return true
}
