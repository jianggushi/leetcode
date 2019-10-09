package solution

// Given a sorted linked list, delete all duplicates such that each element appear only once.
//
// Example 1:
// Input: 1->1->2
// Output: 1->2
//
// Example 2:
// Input: 1->1->2->3->3
// Output: 1->2->3

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

// 只做 value 的移动
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	prev, cur := head, head.Next

	for cur != nil {
		if prev != nil && cur.Val != prev.Val {
			prev = prev.Next
			prev.Val = cur.Val
		}
		cur = cur.Next
	}
	prev.Next = nil
	return head
}

// 一个指针
func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	p := head
	for p.Next != nil {
		if p.Val == p.Next.Val {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}
	}
	return head
}

// 前后两个指针
func deleteDuplicates1(head *ListNode) *ListNode {
	var prev, cur *ListNode = nil, head
	for cur != nil {
		if prev != nil && cur.Val == prev.Val {
			prev.Next = cur.Next
		} else {
			prev = cur
		}
		cur = cur.Next
	}
	return head
}
