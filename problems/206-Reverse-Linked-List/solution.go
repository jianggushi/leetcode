package solution

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 递归方式
// n1 -> ... -> nk-1 -> nk -> nk+1 -> ... -> nm -> $
// 假设 nk+1 到 nm 已经反转，当前需要反转 nk
// n1 -> ... -> nk-1 -> nk -> nk+1 <- ... <- nm
// 那么需要执行
// nk.next.next = nk
// nk.next = nil
func reverseList5(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := reverseList5(head.Next)
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
func reverseList(head *ListNode) *ListNode {
	var l, r *ListNode
	r = head
	for r != nil {
		r.Next, l, r = l, r, r.Next
	}
	return l
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
