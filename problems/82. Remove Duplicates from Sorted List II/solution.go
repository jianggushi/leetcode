package solution

// Given a sorted linked list, delete all nodes that have duplicate numbers, leaving only distinct numbers from the original list.
//
// Example 1:
// Input: 1->2->3->3->4->4->5
// Output: 1->2->5
//
// Example 2:
// Input: 1->1->1->2->3
// Output: 2->3

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
func deleteDuplicates(head *ListNode) *ListNode {
	var prev, cur *ListNode = nil, head

	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			for cur = cur.Next; cur.Next != nil; cur = cur.Next {
				if cur.Val != cur.Next.Val {
					break
				}
			}
			if prev != nil {
				prev.Next = cur.Next
			} else {
				head = cur.Next
			}
		} else {
			prev = cur
		}
		cur = cur.Next
	}
	return head
}
