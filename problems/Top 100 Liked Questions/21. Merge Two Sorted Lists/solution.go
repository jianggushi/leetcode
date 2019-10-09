// Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.

// Example 1:
// Input: 1->2->4, 1->3->4
// Output: 1->1->2->3->4->4

package solution

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
