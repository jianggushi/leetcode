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

func Test_swapPairs_1(t *testing.T) {
	l1 := newListNode([]int{1, 2, 3, 4})
	printListNode(l1)
	printListNode(swapPairs(l1))
}

func Test_swapPairs_2(t *testing.T) {
	l1 := newListNode([]int{1, 2, 3})
	printListNode(l1)
	printListNode(swapPairs(l1))
}
