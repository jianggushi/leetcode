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

func Test_rotateRight_1(t *testing.T) {
	l := newListNode([]int{1, 2, 3, 4, 5})
	printListNode(l)
	printListNode(rotateRight(l, 2))
}

func Test_rotateRight_2(t *testing.T) {
	l := newListNode([]int{0, 1, 2})
	printListNode(l)
	printListNode(rotateRight(l, 4))
}

func Test_rotateRight_3(t *testing.T) {
	l := newListNode([]int{1})
	printListNode(l)
	printListNode(rotateRight(l, 4))
}

func Test_rotateRight_4(t *testing.T) {
	l := newListNode([]int{1, 2})
	printListNode(l)
	printListNode(rotateRight(l, 4))
}

func Test_rotateRight_5(t *testing.T) {
	l := newListNode([]int{1, 2})
	printListNode(l)
	printListNode(rotateRight(l, 1))
}
