package solution

import (
	"reflect"
	"testing"
)

// Given binary tree [3,9,20,null,null,15,7],

//     3
//    / \
//   9  20
//     /  \
//    15   7
func Test_levelOrder1(t *testing.T) {
	root := &TreeNode{
		Val: 3,
	}
	root.Left = &TreeNode{
		Val: 9,
	}
	root.Right = &TreeNode{
		Val: 20,
	}
	root.Right.Left = &TreeNode{
		Val: 15,
	}
	root.Right.Right = &TreeNode{
		Val: 7,
	}
	expected := [][]int{
		{3},
		{9, 20},
		{15, 7},
	}
	ans := levelOrder(root)
	if !reflect.DeepEqual(ans, expected) {
		t.Errorf("expected: %v, levelOrder: %v", expected, ans)
	}
}
