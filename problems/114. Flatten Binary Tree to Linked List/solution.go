package solution

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 解法一
// 递归
// 右 - 左 - 中 遍历方式
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Right)
	flatten(root.Left)
	// 把右子树连接到左子树下，然后替换右子树，左子树置空
	if root.Left != nil {
		node := root.Left
		for node.Right != nil {
			node = node.Right
		}
		node.Right = root.Right
		root.Right = root.Left
		root.Left = nil
	}
}
