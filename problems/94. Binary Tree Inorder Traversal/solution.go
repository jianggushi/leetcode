package solution

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	output := make([]int, 0)
	inorder(root, &output)
	return output
}

func inorder(root *TreeNode, output *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, output)
	*output = append(*output, root.Val)
	inorder(root.Right, output)
}
