package solution

//  Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 快慢指针
// 慢指针每次移动一步，快指针每次移动两步
// 当快指针到达链表末尾，慢指针就到了中间节点
// 奇数个节点，中间节点只有一个
// 偶数个节点，中间节点有两个，这里返回后面的一个
func middleNode2(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// 先遍历一遍链表，统计节点个数
// 再遍历链表，到达中点时停止
func middleNode1(head *ListNode) *ListNode {
	n := 0
	for p := head; p != nil; p = p.Next {
		n++
	}
	m := n / 2
	p := head
	for i := 0; i < m; i++ {
		if i == m {
			break
		}
		p = p.Next
	}
	return p
}
