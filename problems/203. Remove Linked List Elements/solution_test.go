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

func Test_removeElements_1(t *testing.T) {
	l := newListNode([]int{1, 2, 6, 3, 4, 5, 6})
	printListNode(l)
	printListNode(removeElements(l, 6))
}

func Test_removeElements_2(t *testing.T) {
	l := newListNode([]int{1, 2, 6, 3, 4, 5, 6})
	printListNode(l)
	printListNode(removeElements(l, 1))
}

func Test_removeElements_3(t *testing.T) {
	l := newListNode([]int{1})
	printListNode(l)
	printListNode(removeElements(l, 1))
}
