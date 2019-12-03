package solution

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 解法一
// 递归
func buildTree(preorder []int, inorder []int) *TreeNode {
	n := len(inorder)
	if n == 0 {
		return nil
	}
	// preorder 首元素作为当前子树的根节点
	root := &TreeNode{
		Val: preorder[0],
	}
	// 从 inorder 中找到对应的节点元素，并记下位置，以此为分割
	// 左侧是左子树元素，右侧是右子树元素
	pivot := 0
	for i := 0; i < n; i++ {
		if inorder[i] == preorder[0] {
			pivot = i
		}
	}
	// 建立左子树和右子树
	root.Left = buildTree(preorder[1:pivot+1], inorder[:pivot])
	root.Right = buildTree(preorder[pivot+1:], inorder[pivot+1:])
	return root
}
