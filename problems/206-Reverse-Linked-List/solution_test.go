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

func toSlice(node *ListNode) []int {
	res := []int{}
	for ; node != nil; node = node.Next {
		res = append(res, node.Val)
	}
	return res
}

func Test_reverseList_01(t *testing.T) {
	node := new([]int{1, 2, 3, 4, 5})
	expected := []int{5, 4, 3, 2, 1}
	ans := toSlice(reverseList(node))
	assert.Equal(t, expected, ans)
}

func Test_reverseList_02(t *testing.T) {
	node := new([]int{1})
	expected := []int{1}
	ans := toSlice(reverseList(node))
	assert.Equal(t, expected, ans)
}

func Test_reverseList_03(t *testing.T) {
	node := new([]int{1, 2})
	expected := []int{2, 1}
	ans := toSlice(reverseList(node))
	assert.Equal(t, expected, ans)
}

func Test_reverseList_4(t *testing.T) {
	var node *ListNode
	expected := []int{}
	ans := toSlice(reverseList(node))
	assert.Equal(t, expected, ans)
}
