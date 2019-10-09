package solution

// Given a linked list, remove the n-th node from the end of list and return its head.
//
// Example 1:
// Given linked list: 1->2->3->4->5, and n = 2.
// After removing the second node from the end, the linked list becomes 1->2->3->5.
//
// Note:
// Given n will always be valid.

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
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}
	// 遍历链表，计算链表长度
	l := 0
	for p := head; p != nil; p = p.Next {
		l++
	}
	// 待删除节点的顺序编号
	m := l - n
	// 要删除的是头节点
	if m == 0 {
		return head.Next
	}
	// 其他节点
	i := 0
	for p := head; p != nil; p = p.Next {
		i++
		if i == m {
			p.Next = p.Next.Next
		}
	}
	return head
}
