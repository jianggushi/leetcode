// Source : https://leetcode.cn/problems/isomorphic-strings/
// Author : jianggushi
// Date   : 2022-10-27

package solution

import (
	"testing"
)

func Test_isIsomorphic(t *testing.T) {
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
				s: "egg",
				t: "add",
			},
			true,
		},
		{
			"test",
			args{
				s: "foo",
				t: "bar",
			},
			false,
		},
		{
			"test",
			args{
				s: "paper",
				t: "title",
			},
			true,
		},
		{
			"test",
			args{
				s: "badc",
				t: "baba",
			},
			false,
		},
		{
			"test",
			args{
				s: "egcd",
				t: "adfd",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isIsomorphic(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isIsomorphic() = %v, want %v", got, tt.want)
			}
		})
	}
}
