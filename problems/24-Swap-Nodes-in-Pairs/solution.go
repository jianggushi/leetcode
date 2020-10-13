package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	head = &ListNode{Next: head}
	p1 := head
	p2 := head.Next
	for ; p1 != nil && p2 != nil && p2.Next != nil; p1, p2 = p2, p2.Next {
		p1.Next = p2.Next
		p2.Next = p2.Next.Next
		p1.Next.Next = p2
	}
	return head.Next
}
