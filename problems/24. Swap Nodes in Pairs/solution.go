package solution

// Given a linked list, swap every two adjacent nodes and return its head.
// You may not modify the values in the list's nodes, only nodes itself may be changed.
//
// Example 1:
// Given 1->2->3->4, you should return the list as 2->1->4->3.

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
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
