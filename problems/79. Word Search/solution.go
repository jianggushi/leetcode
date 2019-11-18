package solution

func exist(board [][]byte, word string) bool {
	flag := false
	backtrack(board, word, 0, -1, -1, &flag)
	return flag
}

func backtrack(board [][]byte, word string, first int, row int, col int, flag *bool) {
	if first == len(word) {
		*flag = true
		return
	}
	i := 0
	if row-1 > i {
		i = row - 1
	}
	for ; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if !*flag && board[i][j] == word[first] &&
				((row == -1 && col == -1) ||
					(i == row && (j == col-1 || j == col+1)) ||
					(j == col && (i == row-1 || i == row+1))) {
				tmp := board[i][j]
				board[i][j] = 0
				backtrack(board, word, first+1, i, j, flag)
				board[i][j] = tmp
			}
		}
	}
}
