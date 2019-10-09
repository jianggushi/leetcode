package solution

import "testing"

func newTree(values []interface{}) *TreeNode {
	n := len(values)
	if n == 0 {
		return nil
	}
	if values[0] == nil {
		return nil
	}
	queue := make([]*TreeNode, 0)
	root := &TreeNode{}
	index := 0
	root.Val = values[0].(int)
	queue = append(queue, root)
	index++

	for len(queue) != 0 && index < n {
		node := queue[0]
		queue = queue[1:]
		value := values[index]
		index++
		if value != nil {
			lnode := &TreeNode{}
			node.Val = value.(int)
			node.Left = lnode
			queue = append(queue, lnode)
		}
		value = values[index]
		index++
		if value != nil {
			rnode := &TreeNode{}
			node.Val = value.(int)
			node.Right = rnode
			queue = append(queue, rnode)
		}
	}
	return root
}

// 1
// / \
// 2   2
// \   \
// 3    3
func Test_isSymmetric_1(t *testing.T) {
	root := &TreeNode{}
	root.Val = 1

	node := &TreeNode{}
	node.Val = 2
	root.Left = node

	node = &TreeNode{}
	node.Val = 2
	root.Right = node

	node = &TreeNode{}
	node.Val = 3
	root.Left.Right = node

	node = &TreeNode{}
	node.Val = 3
	root.Right.Right = node

	result := isSymmetric(root)

	if result != false {
		t.Error(result)
	}
}
