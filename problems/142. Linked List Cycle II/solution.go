package solution

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	// 快慢指针检测环，得到相遇位置
	slow, fast := head.Next, head.Next.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return nil
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	meet := fast
	// 检测环入口位置
	for head != meet {
		head = head.Next
		meet = meet.Next
	}
	return meet
}
