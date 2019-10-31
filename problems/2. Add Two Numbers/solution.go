package solution

// ListNode definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var l3 *ListNode
	t := 0
	p1, p2, p3 := l1, l2, l3
	for p1 != nil && p2 != nil {
		p3 = new(ListNode)
		s := p1.Val + p2.Val + t
		p3.Val = s % 10
		t = s / 10
		p1 = p1.Next
		p2 = p2.Next
	}
	if p2 != nil {
		return l2
	}
	return l1
}
