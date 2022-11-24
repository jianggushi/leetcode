// Source : https://leetcode.cn/problems/longest-palindrome/
// Author : jianggushi
// Date   : 2022-11-15

package solution

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"test",
			args{
				"abccccdd",
			},
			7,
		},
		{
			"test",
			args{
				"a",
			},
			1,
		},
		{
			"test",
			args{
				"abbbccccaa",
			},
			9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
