package solution

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val > l2.Val {
		l1, l2 = l2, l1
	}

	for p2 := l2; p2 != nil; p2 = l2 {
		l2 = p2.Next
		for p1 := l1; p1 != nil; p1 = p1.Next {
			if (p2.Val >= p1.Val) && (p1.Next == nil || p2.Val < p1.Next.Val) {
				p2.Next = p1.Next
				p1.Next = p2
				break
			}
		}
	}
	return l1
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dumpy := &ListNode{}
	p := dumpy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = l1
			p = p.Next
			l1 = l1.Next
		} else {
			p.Next = l2
			p = p.Next
			l2 = l2.Next
		}
	}
	for l1 != nil {
		p.Next = l1
		p = p.Next
		l1 = l1.Next
	}
	for l2 != nil {
		p.Next = l2
		p = p.Next
		l2 = l2.Next
	}
	return dumpy.Next
}
