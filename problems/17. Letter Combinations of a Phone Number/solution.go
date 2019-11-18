package solution

var maps = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func letterCombinations(digits string) []string {
	ans := make([]string, 0)
	if len(digits) != 0 {
		backtrack(digits, "", &ans)
	}
	return ans
}

func backtrack(digits string, curr string, ans *[]string) {
	if len(digits) == 0 {
		*ans = append(*ans, curr)
		return
	}
	n := len(maps[digits[0]])
	for i := 0; i < n; i++ {
		curr += maps[digits[0]][i : i+1]
		backtrack(digits[1:], curr, ans)
		curr = curr[:len(curr)-1]
	}
}
