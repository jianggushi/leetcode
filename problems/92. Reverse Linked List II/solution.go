package solution

// Reverse a linked list from position m to n. Do it in one-pass.
// Note: 1 ≤ m ≤ n ≤ length of list.
//
// Example:
// Input: 1->2->3->4->5->NULL, m = 2, n = 4
// Output: 1->4->3->2->5->NULL

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
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil {
		return head
	}

	var prev, cur *ListNode = nil, head
	for m > 1 {
		prev, cur = cur, cur.Next
		m--
		n--
	}

	con, tail := prev, cur

	for n > 0 {
		third := cur.Next
		cur.Next = prev
		prev = cur
		cur = third
		n--
	}

	if con != nil {
		con.Next = prev
	} else {
		head = prev
	}
	tail.Next = cur

	return head
}
