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

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	nodeMap := make(map[*ListNode]interface{})
	for headA != nil {
		nodeMap[headA] = nil
		headA = headA.Next
	}
	for headB != nil {
		if _, ok := nodeMap[headB]; ok {
			return headB
		}
		headB = headA.Next
	}
	return nil
}

func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	for pa := headA; pa != nil; pa = pa.Next {
		for pb := headB; pb != nil; pb = pb.Next {
			if pa == pb {
				return pa
			}
		}
	}
	return nil
}
