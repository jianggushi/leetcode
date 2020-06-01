package solution

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 使用递归遍历
// 回溯算法
// 递归时，带上当前的层次 level
// 把当前 level 的节点加到对应的 level 数组中
func levelOrder(root *TreeNode) [][]int {
	ans := [][]int{}
	helper(root, 0, &ans)
	return ans
}

func helper(root *TreeNode, level int, ans *[][]int) {
	if root == nil {
		return
	}
	// create a new level
	if level == len(*ans) {
		*ans = append(*ans, []int{})
	}
	(*ans)[level] = append((*ans)[level], root.Val)
	helper(root.Left, level+1, ans)
	helper(root.Right, level+1, ans)
}

// 每次遍历一层
func levelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	ans := make([][]int, 0)
	q := []*TreeNode{root}
	for i := 0; len(q) != 0; i++ {
		ans = append(ans, []int{})
		// 临时队列 p 暂存下一层节点
		p := []*TreeNode{}
		// 遍历完当前层所有的节点
		for j := 0; j < len(q); j++ {
			node := q[j]
			ans[i] = append(ans[i], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return ans
}

// 使用二元组 (node, level) 记录节点和它所在的层数
// 每个新入队的节点的 level 等于它的父节点的 level + 1
func levelOrder1(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	ans := make([][]int, 0)
	type TreeNodeLevel struct {
		node  *TreeNode
		level int
	}
	q := []TreeNodeLevel{
		{
			node:  root,
			level: 0,
		},
	}
	for len(q) != 0 {
		e := q[0]
		q = q[1:]
		// create a new level
		if e.level == len(ans) {
			ans = append(ans, []int{})
		}
		ans[e.level] = append(ans[e.level], e.node.Val)
		if e.node.Left != nil {
			q = append(q, TreeNodeLevel{node: e.node.Left, level: e.level + 1})
		}
		if e.node.Right != nil {
			q = append(q, TreeNodeLevel{node: e.node.Right, level: e.level + 1})
		}
	}
	return ans
}
