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

func Test_partition_1(t *testing.T) {
	l := newListNode([]int{1, 4, 3, 2, 5, 2})
	printListNode(l)
	printListNode(partition(l, 3))
}

func Test_partition_2(t *testing.T) {
	l := newListNode([]int{1, 4, 3, 2, 5, 2})
	printListNode(l)
	printListNode(partition(l, 1))
}

func Test_partition_3(t *testing.T) {
	l := newListNode([]int{1, 4, 3, 2, 5, 2})
	printListNode(l)
	printListNode(partition(l, 6))
}

func Test_partition_4(t *testing.T) {
	l := newListNode([]int{2, 1})
	printListNode(l)
	printListNode(partition(l, 2))
}

func Test_partition_5(t *testing.T) {
	l := newListNode([]int{2, 1, 3})
	printListNode(l)
	printListNode(partition(l, 2))
}
