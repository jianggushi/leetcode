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

func Test_reverseBetween_1(t *testing.T) {
	l := newListNode([]int{1, 2, 3, 4, 5})
	printListNode(l)
	printListNode(reverseBetween(l, 2, 4))
}

func Test_reverseBetween_2(t *testing.T) {
	l := newListNode([]int{1, 2, 3, 4, 5})
	printListNode(l)
	printListNode(reverseBetween(l, 1, 4))
}

func Test_reverseBetween_3(t *testing.T) {
	l := newListNode([]int{1, 2, 3, 4, 5})
	printListNode(l)
	printListNode(reverseBetween(l, 2, 5))
}

func Test_reverseBetween_4(t *testing.T) {
	l := newListNode([]int{1, 2, 3, 4, 5})
	printListNode(l)
	printListNode(reverseBetween(l, 1, 5))
}

func Test_reverseBetween_5(t *testing.T) {
	l := newListNode([]int{1, 2, 3, 4, 5})
	printListNode(l)
	printListNode(reverseBetween(l, 1, 1))
}

func Test_reverseBetween_6(t *testing.T) {
	l := newListNode([]int{1, 2, 3, 4, 5})
	printListNode(l)
	printListNode(reverseBetween(l, 2, 2))
}

func Test_reverseBetween_7(t *testing.T) {
	l := newListNode([]int{1, 2})
	printListNode(l)
	printListNode(reverseBetween(l, 1, 2))
}

func Test_reverseBetween_8(t *testing.T) {
	l := newListNode([]int{1})
	printListNode(l)
	printListNode(reverseBetween(l, 1, 1))
}
