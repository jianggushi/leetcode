package solution

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 只遍历一遍
// 为了兼容处理删除头结点的情况，创建一个 dumpy 节点
// 使用两个指针 left 和 right
// 先移动 right 使 right 和 left 相隔 n
// 然后再迭代链表，分别向后移动 left 和 right
// right 移动到尾结点时，删除 left 后一个节点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dumpy := &ListNode{
		Next: head,
	}
	l, r := dumpy, dumpy
	for i := 0; i < n; i++ {
		r = r.Next
	}
	for r != nil && r.Next != nil {
		l = l.Next
		r = r.Next
	}
	l.Next = l.Next.Next
	return dumpy.Next
}

// 遍历计算链表长度 l
// 要删除的节点是第 l-n 个
// 找到第 l-n-1 个节点，删除它后面的节点
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	l := 0
	for p := head; p != nil; p = p.Next {
		l++
	}
	if l-n == 0 {
		return head.Next
	}
	i := 1
	for p := head; p != nil; p = p.Next {
		if i == l-n && p.Next != nil {
			p.Next = p.Next.Next
			break
		}
		i++
	}
	return head
}
