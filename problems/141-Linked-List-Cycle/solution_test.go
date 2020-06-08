package solution

import (
	"fmt"
	"testing"
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

func Test_hasCycle_1(t *testing.T) {
	l := new([]int{1, 2, 6, 3, 4, 5, 6})
	fmt.Println(hasCycle(l))
}
