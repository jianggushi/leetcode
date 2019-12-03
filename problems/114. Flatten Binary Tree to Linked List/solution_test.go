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

func Test_flatten_1(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 5,
			Right: &TreeNode{
				Val: 6,
			},
		},
	}
	expected := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 4,
					Right: &TreeNode{
						Val: 5,
						Right: &TreeNode{
							Val: 6,
						},
					},
				},
			},
		},
	}
	flatten(root)
	if !EqualTree(root, expected) {
		t.Errorf("expected: %v, flatten: %v", expected, root)
	}
}
