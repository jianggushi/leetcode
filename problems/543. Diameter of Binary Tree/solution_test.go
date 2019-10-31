package solution

import (
	"testing"
)

func Test_diameterOfBinaryTree_1(t *testing.T) {
	root := &TreeNode{Val: 1}
	node := &TreeNode{Val: 2}
	root.Left = node
	node = &TreeNode{Val: 3}
	root.Right = node

	t.Error(diameterOfBinaryTree(root))
}
