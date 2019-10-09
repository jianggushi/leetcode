package solution

// Given a linked list and a value x, partition it such that all nodes less than x come before nodes greater than or equal to x.
// You should preserve the original relative order of the nodes in each of the two partitions.
//
// Example:
// Input: head = 1->4->3->2->5->2, x = 3
// Output: 1->2->2->4->3->5

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
func partition(head *ListNode, x int) *ListNode {
	var l, lp, r, rp *ListNode
	for head != nil {
		if head.Val < x {
			if l == nil {
				l = head
			} else {
				lp.Next = head
			}
			lp = head
			head = head.Next
			lp.Next = nil
		} else {
			if r == nil {
				r = head
			} else {
				rp.Next = head
			}
			rp = head
			head = head.Next
			rp.Next = nil
		}
	}
	// 合并
	if l == nil {
		l = r
	} else {
		lp.Next = r
	}
	return l
}
