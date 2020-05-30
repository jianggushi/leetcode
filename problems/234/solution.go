package solution

// ListNode definition for singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// 方法一：暴力扫描
// 定义两个指针 pr 和 pl，分别从链表左边和右边扫描，
// 比较 pr 和 pl 的 val 是否相等，如果不相等，那么返回 false。
// 因为是单链表，pl 很容易确定，从左边开始扫描就可以，
// 但是 pr 怎么确定呢？只能每次从链表左边开始扫描，每次找到 pr 的上一个节点。
// 终止条件呢？如果是奇数节点，那么最后 pl == pr，如果是偶数节点，那么最后 pl.Next == pr。
// 时间复杂度呢？因为每次都要从左到右找到 pr 的位置，平均时间复杂度 O(n2)

// 方法二：找到中间节点，逆序一半链表，然后从两段链表的两端进行回文检查

// 找到中间节点，逆序前一半，在一个循环中进行
func isPalindrome6(head *ListNode) bool {
	if head == nil {
		return true
	}
	// 快慢指针找到中间节点
	slow, fast := head, head
	var prev *ListNode
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		next := slow.Next
		slow.Next = prev
		prev = slow
		slow = next
	}
	// 如果有奇数节点，slow 指向中间节点，prev 指向前面一个阶段
	// 1 - 2 -> 3 -> 2 -> 1
	//     |    |
	//    prev slow
	if fast != nil {
		slow = slow.Next
	}
	// 检测是否回文链表
	for prev != nil {
		if prev.Val != slow.Val {
			return false
		}
		prev = prev.Next
		slow = slow.Next
	}
	return true
}

// 找到中间节点，逆序前一半
func isPalindrome5(head *ListNode) bool {
	if head == nil {
		return true
	}
	// 快慢指针找到中间节点
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 逆序前一半
	var prev, curr, next *ListNode
	curr = head
	for curr != slow {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	// 如果有奇数节点，slow 指向中间节点，prev 指向前面一个阶段
	// 1 - 2 -> 3 -> 2 -> 1
	//     |    |
	//    prev slow
	if fast != nil {
		slow = slow.Next
	}
	// 检测是否回文链表
	for prev != nil {
		if prev.Val != slow.Val {
			return false
		}
		prev = prev.Next
		slow = slow.Next
	}
	return true
}

// 找到中间节点，逆序后一半
func isPalindrome4(head *ListNode) bool {
	if head == nil {
		return true
	}
	// 快慢指针找到中间节点
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 逆序后一半
	p1, p2 := slow, slow.Next
	p1.Next = nil
	for p2 != nil {
		tmp := p2.Next
		p2.Next = p1
		p1 = p2
		p2 = tmp
	}
	// 检测是否回文链表
	pr := p1
	pl := head
	for pr != nil {
		if pl.Val != pr.Val {
			return false
		}
		pl = pl.Next
		pr = pr.Next
	}
	return true
}

// 找到中间节点，逆序后一半
func isPalindrome3(head *ListNode) bool {
	if head == nil {
		return true
	}
	// 找到中间节点
	n := 0
	for p := head; p.Next != nil; p = p.Next {
		n++
	}
	p1 := head
	for i := 1; i < (n+1)/2; i++ {
		p1 = p1.Next
	}
	// 逆序后一半
	p2 := p1.Next
	p1.Next = nil
	for p2 != nil {
		tmp := p2.Next
		p2.Next = p1
		p1 = p2
		p2 = tmp
	}
	// 检测是否回文链表
	pr := p1
	pl := head
	for pl != nil {
		if pl.Val != pr.Val {
			return false
		}
		pl = pl.Next
		pr = pr.Next
	}
	return true
}

// 暴力扫描
func isPalindrome(head *ListNode) bool {
	var tail, l, r *ListNode
	l = head
	for l != tail && l.Next != r {
		for r = l; r.Next != tail; r = r.Next {
		}
		if l.Val != r.Val {
			return false
		}
		tail = r
		l = l.Next
	}
	return true
}
