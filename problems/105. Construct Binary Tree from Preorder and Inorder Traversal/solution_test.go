package solution

import "testing"

func EqualTree(r1 *TreeNode, r2 *TreeNode) bool {
	if r1 == nil && r2 == nil {
		return true
	}
	if r1 == nil && r2 != nil {
		return false
	}
	if r1 != nil && r2 == nil {
		return false
	}
	if r1.Val != r2.Val {
		return false
	}
	if !EqualTree(r1.Left, r2.Left) {
		return false
	}
	if !EqualTree(r1.Right, r2.Right) {
		return false
	}
	return true
}

func Test_buildTree_1(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	expected := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	ans := buildTree(preorder, inorder)
	if !EqualTree(ans, expected) {
		t.Errorf("expected: %v, buildTree: %v", expected, ans)
	}
}
