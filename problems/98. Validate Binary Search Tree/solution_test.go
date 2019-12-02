package solution

import "testing"

func Test_isValidBST_1(t *testing.T) {
	root := &TreeNode{
		Val: 2,
	}
	root.Left = &TreeNode{
		Val: 1,
	}
	root.Right = &TreeNode{
		Val: 3,
	}
	expected := true
	ans := isValidBST(root)
	if ans != expected {
		t.Errorf("expected: %v, isValidBST: %v", expected, ans)
	}
}

func Test_isValidBST_2(t *testing.T) {
	root := &TreeNode{
		Val: 5,
	}
	root.Left = &TreeNode{
		Val: 1,
	}
	root.Right = &TreeNode{
		Val: 4,
	}
	root.Right.Left = &TreeNode{
		Val: 3,
	}
	root.Right.Right = &TreeNode{
		Val: 6,
	}
	expected := false
	ans := isValidBST(root)
	if ans != expected {
		t.Errorf("expected: %v, isValidBST: %v", expected, ans)
	}
}

func Test_isValidBST_3(t *testing.T) {
	root := &TreeNode{
		Val: 10,
	}
	root.Left = &TreeNode{
		Val: 5,
	}
	root.Right = &TreeNode{
		Val: 15,
	}
	root.Right.Left = &TreeNode{
		Val: 6,
	}
	root.Right.Right = &TreeNode{
		Val: 20,
	}
	expected := false
	ans := isValidBST(root)
	if ans != expected {
		t.Errorf("expected: %v, isValidBST: %v", expected, ans)
	}
}
