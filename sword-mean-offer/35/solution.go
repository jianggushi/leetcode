package solution

// Definition for a Node.
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	map1 := make(map[*Node]int)
	map2 := make(map[int]*Node)
	ans := &Node{}
	i := 0
	p1, p2 := head, ans
	for p1 != nil {
		node := &Node{
			Val: p1.Val,
		}
		map1[p1] = i
		map2[i] = node
		p1 = p1.Next
		p2.Next = node
		p2 = p2.Next
		i++
	}
	p1, p2 = head, ans.Next
	for p1 != nil {
		if p1.Random != nil {
			p2.Random = map2[map1[p1.Random]]
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	return ans.Next
}
