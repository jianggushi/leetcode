package solution

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 先暂存 curr.Next
// 最终 curr = nil, prev 变成了 head 节点
func reverseList(head *ListNode) *ListNode {
	var prev, curr *ListNode
	curr = head
	for curr != nil {
		tmp := curr.Next
		curr.Next = prev
		prev = curr
		curr = tmp
	}
	return prev
}
