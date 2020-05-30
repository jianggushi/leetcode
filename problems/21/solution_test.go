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

func Test_mergeTwoLists_01(t *testing.T) {
	l1 := new([]int{1, 2, 4})
	l2 := new([]int{1, 3, 4})
	expected := []int{1, 1, 2, 3, 4, 4}
	ans := toSlice(mergeTwoLists(l1, l2))
	assert.Equal(t, expected, ans)
}

func Test_mergeTwoLists_02(t *testing.T) {
	var l1, l2 *ListNode
	expected := []int{}
	ans := toSlice(mergeTwoLists(l1, l2))
	assert.Equal(t, expected, ans)
}

func Test_mergeTwoLists_03(t *testing.T) {
	var l1 *ListNode
	l2 := new([]int{1, 3, 4})
	expected := []int{1, 3, 4}
	ans := toSlice(mergeTwoLists(l1, l2))
	assert.Equal(t, expected, ans)
}

func Test_mergeTwoLists_04(t *testing.T) {
	l1 := new([]int{1, 2, 4})
	var l2 *ListNode
	expected := []int{1, 2, 4}
	ans := toSlice(mergeTwoLists(l1, l2))
	assert.Equal(t, expected, ans)
}

func Test_mergeTwoLists_05(t *testing.T) {
	l1 := new([]int{1, 2, 3, 4})
	l2 := new([]int{5, 6, 7, 8})
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8}
	ans := toSlice(mergeTwoLists(l1, l2))
	assert.Equal(t, expected, ans)
}

func Test_mergeTwoLists_06(t *testing.T) {
	l1 := new([]int{5, 6, 7, 8})
	l2 := new([]int{1, 2, 3, 4})
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8}
	ans := toSlice(mergeTwoLists(l1, l2))
	assert.Equal(t, expected, ans)
}

func Test_mergeTwoLists_07(t *testing.T) {
	l1 := new([]int{5, 6, 7, 8})
	l2 := new([]int{5, 6, 7, 8})
	expected := []int{5, 5, 6, 6, 7, 7, 8, 8}
	ans := toSlice(mergeTwoLists(l1, l2))
	assert.Equal(t, expected, ans)
}

func Test_mergeTwoLists_08(t *testing.T) {
	l1 := new([]int{5, 5, 5, 5})
	l2 := new([]int{5, 5, 5, 5})
	expected := []int{5, 5, 5, 5, 5, 5, 5, 5}
	ans := toSlice(mergeTwoLists(l1, l2))
	assert.Equal(t, expected, ans)
}

func Test_mergeTwoLists_09(t *testing.T) {
	l1 := new([]int{4, 4, 4, 4})
	l2 := new([]int{5, 5, 5, 5})
	expected := []int{4, 4, 4, 4, 5, 5, 5, 5}
	ans := toSlice(mergeTwoLists(l1, l2))
	assert.Equal(t, expected, ans)
}
