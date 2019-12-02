package solution

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return backtrack(root, nil, nil)
}

func backtrack(root *TreeNode, lower *int, upper *int) bool {
	if root == nil {
		return true
	}
	// left sub tree
	if upper != nil && root.Val >= *upper {
		return false
	}
	// right sub tree
	if lower != nil && root.Val <= *lower {
		return false
	}
	// left sub tree
	if !backtrack(root.Left, lower, &root.Val) {
		return false
	}
	// right sub tree
	if !backtrack(root.Right, &root.Val, upper) {
		return false
	}
	return true
}
