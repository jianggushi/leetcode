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

func Test_removeNthFromEnd_01(t *testing.T) {
	node := new([]int{1, 2, 3, 4, 5})
	n := 2
	expected := []int{1, 2, 3, 5}
	ans := toSlice(removeNthFromEnd(node, n))
	assert.Equal(t, expected, ans)
}

func Test_removeNthFromEnd_02(t *testing.T) {
	node := new([]int{1, 2, 3, 4, 5})
	n := 1
	expected := []int{1, 2, 3, 4}
	ans := toSlice(removeNthFromEnd(node, n))
	assert.Equal(t, expected, ans)
}

func Test_removeNthFromEnd_03(t *testing.T) {
	node := new([]int{1, 2, 3, 4, 5})
	n := 5
	expected := []int{2, 3, 4, 5}
	ans := toSlice(removeNthFromEnd(node, n))
	assert.Equal(t, expected, ans)
}

func Test_removeNthFromEnd_04(t *testing.T) {
	node := new([]int{1, 2, 3, 4, 5})
	n := 4
	expected := []int{1, 3, 4, 5}
	ans := toSlice(removeNthFromEnd(node, n))
	assert.Equal(t, expected, ans)
}

func Test_removeNthFromEnd_05(t *testing.T) {
	node := new([]int{1})
	n := 1
	expected := []int{}
	ans := toSlice(removeNthFromEnd(node, n))
	assert.Equal(t, expected, ans)
}
