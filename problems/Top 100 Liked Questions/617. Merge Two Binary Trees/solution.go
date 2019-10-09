package solution

// TreeNode definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	t1.Val += t2.Val
	t1.Left = mergeTrees(t1.Left, t2.Left)
	t1.Right = mergeTrees(t1.Right, t2.Right)
	return t1
}

func mergeTrees1(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	}
	if t1 != nil && t2 == nil {
		return t1
	}
	if t1 == nil && t2 != nil {
		return t2
	}
	if t1 != nil && t2 != nil {
		t1.Val += t2.Val
		if t1.Left == nil && t2.Left != nil {
			t1.Left = t2.Left
		} else {
			mergeTrees(t1.Left, t2.Left)
		}
		if t1.Right == nil && t2.Right != nil {
			t1.Right = t2.Right
		} else {
			mergeTrees(t1.Right, t2.Right)
		}
	}
	return t1
}
