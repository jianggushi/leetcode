package solution

// TreeNode definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	ans := 0
	depth(root, &ans)
	return ans
}

func depth(root *TreeNode, ans *int) int {
	if root == nil {
		return 0
	}
	lh := depth(root.Left, ans)
	rh := depth(root.Right, ans)
	if lh+rh > *ans {
		*ans = lh + rh
	}
	if lh > rh {
		return lh + 1
	}
	return rh + 1
}
