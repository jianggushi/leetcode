package solution

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归遍历，深度优先搜索
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	dl, dr := maxDepth(root.Left), maxDepth(root.Right)
	if dl > dr {
		return dl + 1
	}
	return dr + 1
}
