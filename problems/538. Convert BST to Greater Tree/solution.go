package solution

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 解法一
// 反序中序遍历
func convertBST(root *TreeNode) *TreeNode {
	n := 0
	convert(root, &n)
	return root
}

func convert(root *TreeNode, n *int) {
	if root == nil {
		return
	}
	convert(root.Right, n)
	root.Val += *n
	*n = root.Val
	convert(root.Left, n)
}
