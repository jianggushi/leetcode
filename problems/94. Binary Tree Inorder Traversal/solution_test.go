package solution

import "testing"

func Test_inorderTraversal(t *testing.T) {
	node := &TreeNode{}
	node.Val = 1
	node.Left = nil
	node.Right = &TreeNode{}

	node.Right.Val = 2
	node.Right.Left = &TreeNode{}
	node.Right.Right = nil

	node.Right.Left.Val = 3
	node.Right.Left.Left = nil
	node.Right.Left.Right = nil

	output := inorderTraversal(node)

	t.Error(output)
}
