// Source : https://leetcode.cn/problems/is-subsequence/
// Author : jianggushi
// Date   : 2022-11-02

package solution

import "testing"

func Test_isSubsequence(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			"test",
			args{
				s: "ace",
				t: "abcde",
			},
			true,
		},
		{
			"test",
			args{
				s: "aec",
				t: "abcde",
			},
			false,
		},
		{
			"test",
			args{
				s: "abc",
				t: "ahbgdc",
			},
			true,
		},
		{
			"test",
			args{
				s: "axc",
				t: "ahbgdc",
			},
			false,
		},
		{
			"test",
			args{
				s: "b",
				t: "c",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubsequence(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
