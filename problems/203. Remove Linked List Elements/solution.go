package solution

// Remove all elements from a linked list of integers that have value val.
//
// Example:
// Input:  1->2->6->3->4->5->6, val = 6
// Output: 1->2->3->4->5

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

// 在 removeElements3 的基础上，再减少一个变量名
func removeElements(head *ListNode, val int) *ListNode {
	head = &ListNode{Next: head}
	p := head
	for p != nil && p.Next != nil {
		if p.Next.Val == val {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}
	}
	return head.Next
}

// 现在链表头部增加一个假的节点，然后使用一个指针处理
func removeElements3(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}
	p := dummy
	for p != nil && p.Next != nil {
		if p.Next.Val == val {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}
	}
	return dummy.Next
}

// 先处理 head，让 head 不等于 val，然后使用一个指针处理
func removeElements2(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	for head != nil && head.Val == val {
		head = head.Next
	}

	p := head
	for p != nil && p.Next != nil {
		if p.Next.Val == val {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}
	}
	return head
}

// 前后两个指针，统一逻辑
func removeElements1(head *ListNode, val int) *ListNode {
	var prev, cur *ListNode = nil, head
	for cur != nil {
		if cur.Val == val {
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
