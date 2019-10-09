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

func Test_mergeKLists_1(t *testing.T) {
	l1 := newListNode([]int{1, 4, 5})
	l2 := newListNode([]int{1, 3, 4})
	l3 := newListNode([]int{2, 6})
	printListNode(l1)
	printListNode(l2)
	printListNode(l3)
	printListNode(mergeKLists([]*ListNode{l1, l2, l3}))
}

func Test_mergeKLists_2(t *testing.T) {
	printListNode(mergeKLists(nil))
}

func Test_mergeKLists_3(t *testing.T) {
	printListNode(mergeKLists([]*ListNode{}))
}
