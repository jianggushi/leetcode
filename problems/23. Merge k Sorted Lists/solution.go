package solution

// Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity.
//
// Example 1:
// Input:
// [
//   1->4->5,
// 	 1->3->4,
// 	 2->6
// ]
// Output: 1->1->2->3->4->4->5->6

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
func mergeKLists(lists []*ListNode) *ListNode {
	l := len(lists)
	switch l {
	case 0:
		return nil
	case 1:
		return lists[0]
	case 2:
		return mergeTwoLists(lists[0], lists[1])
	default:
		return mergeTwoLists(mergeKLists(lists[:l/2]), mergeKLists(lists[l/2:]))
	}
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val > l2.Val {
		l1, l2 = l2, l1
	}

	for p2 := l2; p2 != nil; p2 = l2 {
		l2 = p2.Next
		for p1 := l1; p1 != nil; p1 = p1.Next {
			if (p2.Val >= p1.Val) && (p1.Next == nil || p2.Val < p1.Next.Val) {
				p2.Next = p1.Next
				p1.Next = p2
				break
			}
		}
	}
	return l1
}

func mergeKLists1(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
		return nil
	}
	t := lists[0]
	for i := 1; i < len(lists); i++ {
		t = mergeTwoLists(t, lists[i])
	}
	return t
}
