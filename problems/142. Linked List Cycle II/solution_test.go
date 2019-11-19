package solution

import "testing"

func NewListNode(vals []int, pos int) *ListNode {
	head := &ListNode{Val: vals[0]}
	p := head
	for _, v := range vals[1:] {
		t := &ListNode{Val: v}
		p.Next = t
		p = t
	}
	p, tail := head, p
	i := 0
	for p != nil {
		if i == pos {
			tail.Next = p
			break
		}
		p = p.Next
		i++
	}
	return head
}

func Test_detectCycle1(t *testing.T) {
	head := NewListNode([]int{3, 2, 0, -4}, 1)
	expected := head.Next
	ans := detectCycle(head)
	if ans != expected {
		t.Errorf("expected: %v, detectCycle: %v", expected, ans)
	}
}

func Test_detectCycle2(t *testing.T) {
	head := NewListNode([]int{1, 2}, 0)
	expected := head
	ans := detectCycle(head)
	if ans != expected {
		t.Errorf("expected: %v, detectCycle: %v", expected, ans)
	}
}

func Test_detectCycle3(t *testing.T) {
	head := NewListNode([]int{1}, -1)
	var expected *ListNode
	ans := detectCycle(head)
	if ans != expected {
		t.Errorf("expected: %v, detectCycle: %v", expected, ans)
	}
}
