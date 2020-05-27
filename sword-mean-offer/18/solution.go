package solution

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 增加 dummy 节点
func deleteNode(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}
	for p := dummy; p.Next != nil; p = p.Next {
		if p.Next.Val == val {
			p.Next = p.Next.Next
			break
		}
	}
	return dummy.Next
}

// head 节点特殊处理
func deleteNode1(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Val == val {
		return head.Next
	}

	for p := head; p.Next != nil; p = p.Next {
		if p.Next.Val == val {
			p.Next = p.Next.Next
			break
		}
	}
	return head
}
