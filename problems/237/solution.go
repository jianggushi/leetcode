package solution

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 没有给链表头，只给了当前要删除的节点
// 只能替换元素了，然后删除尾巴上的节点
func deleteNode(node *ListNode) {
	var p *ListNode
	for node != nil && node.Next != nil {
		node.Val = node.Next.Val
		p = node
		node = node.Next
	}
	if p != nil {
		p.Next = nil
	}
}
