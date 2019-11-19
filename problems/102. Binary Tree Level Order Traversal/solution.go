package solution

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	ans := make([][]int, 0)
	type TreeNodeLevel struct {
		node  *TreeNode
		level int
	}
	nodes := make([]*TreeNodeLevel, 1)
	nodes[0] = &TreeNodeLevel{
		node:  root,
		level: 0,
	}
	for len(nodes) != 0 {
		node := nodes[0]
		nodes = nodes[1:]
		if node.level == len(ans) {
			ans = append(ans, []int{})
		}
		ans[node.level] = append(ans[node.level], node.node.Val)
		if node.node.Left != nil {
			nodes = append(nodes, &TreeNodeLevel{node: node.node.Left, level: node.level + 1})
		}
		if node.node.Right != nil {
			nodes = append(nodes, &TreeNodeLevel{node: node.node.Right, level: node.level + 1})
		}
	}
	return ans
}
