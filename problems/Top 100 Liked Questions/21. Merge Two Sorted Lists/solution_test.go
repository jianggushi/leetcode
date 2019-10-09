package solution

import (
	"fmt"
	"testing"
)

func newListNode(vals []int) *ListNode {
	l := &ListNode{Val: vals[0]}
	p := l
	for _, v := range vals[1:] {
		t := &ListNode{Val: v}
		p.Next = t
		p = t
	}
	return l
}

func printListNode(l *ListNode) {
	for ; l != nil; l = l.Next {
		fmt.Printf("%v ", l.Val)
	}
	fmt.Println()
}

func Test_mergeTwoLists_1(t *testing.T) {
	l1 := newListNode([]int{1, 2, 4})
	l2 := newListNode([]int{1, 3, 4})
	printListNode(l1)
	printListNode(l2)
	printListNode(mergeTwoLists(l1, l2))
}

func Test_mergeTwoLists_2(t *testing.T) {
	var l1, l2 *ListNode
	printListNode(l1)
	printListNode(l2)
	printListNode(mergeTwoLists(l1, l2))
}

func Test_mergeTwoLists_3(t *testing.T) {
	var l1 *ListNode
	l2 := newListNode([]int{1, 3, 4})
	printListNode(l1)
	printListNode(l2)
	printListNode(mergeTwoLists(l1, l2))
}

func Test_mergeTwoLists_4(t *testing.T) {
	l1 := newListNode([]int{1, 2, 4})
	var l2 *ListNode
	printListNode(l1)
	printListNode(l2)
	printListNode(mergeTwoLists(l1, l2))
}

func Test_mergeTwoLists_5(t *testing.T) {
	l1 := newListNode([]int{1, 2, 3, 4})
	l2 := newListNode([]int{5, 6, 7, 8})
	printListNode(l1)
	printListNode(l2)
	printListNode(mergeTwoLists(l1, l2))
}

func Test_mergeTwoLists_6(t *testing.T) {
	l1 := newListNode([]int{5, 6, 7, 8})
	l2 := newListNode([]int{1, 2, 3, 4})
	printListNode(l1)
	printListNode(l2)
	printListNode(mergeTwoLists(l1, l2))
}

func Test_mergeTwoLists_7(t *testing.T) {
	l1 := newListNode([]int{5, 6, 7, 8})
	l2 := newListNode([]int{5, 6, 7, 8})
	printListNode(l1)
	printListNode(l2)
	printListNode(mergeTwoLists(l1, l2))
}

func Test_mergeTwoLists_8(t *testing.T) {
	l1 := newListNode([]int{5, 5, 5, 5})
	l2 := newListNode([]int{5, 5, 5, 5})
	printListNode(l1)
	printListNode(l2)
	printListNode(mergeTwoLists(l1, l2))
}

func Test_mergeTwoLists_9(t *testing.T) {
	l1 := newListNode([]int{4, 4, 4, 4})
	l2 := newListNode([]int{5, 5, 5, 5})
	printListNode(l1)
	printListNode(l2)
	printListNode(mergeTwoLists(l1, l2))
}
