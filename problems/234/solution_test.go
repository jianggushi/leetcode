package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func new(vals []int) *ListNode {
	l := &ListNode{Val: vals[0]}
	p := l
	for _, v := range vals[1:] {
		t := &ListNode{Val: v}
		p.Next = t
		p = t
	}
	return l
}

func Test_isPalindrome_01(t *testing.T) {
	l := new([]int{1, 2})
	expected := false
	ans := isPalindrome(l)
	assert.Equal(t, expected, ans)
}

func Test_isPalindrome_02(t *testing.T) {
	l := new([]int{1, 2, 2, 1})
	expected := true
	ans := isPalindrome(l)
	assert.Equal(t, expected, ans)
}

func Test_isPalindrome_03(t *testing.T) {
	expected := true
	ans := isPalindrome(nil)
	assert.Equal(t, expected, ans)
}

func Test_isPalindrome_04(t *testing.T) {
	l := new([]int{1})
	expected := true
	ans := isPalindrome(l)
	assert.Equal(t, expected, ans)
}
