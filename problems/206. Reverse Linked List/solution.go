package solution

// Reverse a singly linked list.
//
// Example 1:
// Input: 1->2->3->4->5->NULL
// Output: 5->4->3->2->1->NULL
//
// Follow up:
// A linked list can be reversed either iteratively or recursively. Could you implement both?

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

// 递归方式
// n1 -> ... -> nk-1 -> nk -> nk+1 -> ... -> nm -> $
// 假设 nk+1 到 nm 已经反转，当前需要反转 nk
// n1 -> ... -> nk-1 -> nk -> nk+1 <- ... <- nm
// 那么需要执行
// nk.next.next = nk
// nk.next = nil
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}

// 插入到头部 和 reverseList3 相同
func reverseList4(head *ListNode) *ListNode {
	var last *ListNode

	for head != nil {
		t := head.Next
		head.Next = last
		last = head
		head = t
	}

	return last
}

// 插入到头部
func reverseList3(head *ListNode) *ListNode {
	var last *ListNode

	for head != nil {
		head.Next, head, last = last, head.Next, head
	}

	return last
}

// 两个指针
func reverseList2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	p1 := head
	p2 := head.Next
	for p1 != nil && p2 != nil {
		p1.Next = p2.Next
		p2.Next = head
		head = p2
		p2 = p1.Next
	}
	return head
}

// head 不动
func reverseList1(head *ListNode) *ListNode {
	p := head
	for p != nil && p.Next != nil {
		t := p.Next
		p.Next = p.Next.Next
		t.Next = head
		head = t
	}
	return head
}
