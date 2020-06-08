package solution

// ListNode Definition for singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// 创建一个冗余的头节点，简化头结点处理
func hasCycle(head *ListNode) bool {
	dumpy := &ListNode{
		Next: head,
	}
	slow, fast := dumpy, dumpy.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

// 快慢指针
// 慢指针每次移动一步，快指针每次移动两步
// 如果有环，快指针肯定会追上慢指针
// 如果没有环，快指针和慢指针都会到达末尾
func hasCycle2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

// 哈希表
// 遍历链表节点，使用哈希表保存访问过的节点
// 如果当前节点之前访问过，说明链表中有环
func hasCycle1(head *ListNode) bool {
	mmap := make(map[*ListNode]int)
	for p := head; p != nil; p = p.Next {
		if mmap[p] > 0 {
			return true
		}
		mmap[p]++
	}
	return false
}
