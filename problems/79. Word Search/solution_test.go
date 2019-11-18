package solution

import "testing"

func Test_exist_1(t *testing.T) {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCCED"
	expected := true
	ans := exist(board, word)
	if ans != expected {
		t.Errorf("expected: %v, exist: %v", expected, ans)
	}
}

func Test_exist_2(t *testing.T) {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "SEE"
	expected := true
	ans := exist(board, word)
	if ans != expected {
		t.Errorf("expected: %v, exist: %v", expected, ans)
	}
}

func Test_exist_3(t *testing.T) {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCB"
	expected := false
	ans := exist(board, word)
	if ans != expected {
		t.Errorf("expected: %v, exist: %v", expected, ans)
	}
}
