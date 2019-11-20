package solution

import "testing"

func NewTree(vals []interface{}) *TreeNode {
	n := len(vals)
	if n == 0 {
		return nil
	}
	if vals[0] == nil {
		return nil
	}
	root := &TreeNode{
		Val: vals[0].(int),
	}
	nodes := make([]*TreeNode, 1)
	nodes[0] = root
	i := 1
	for len(nodes) != 0 {
		node := nodes[0]
		nodes = nodes[1:]
		if i < n && vals[i] != nil {
			node.Left = &TreeNode{
				Val: vals[i].(int),
			}
			nodes = append(nodes, node.Left)
		}
		i++
		if i < n && vals[i] != nil {
			node.Right = &TreeNode{
				Val: vals[i].(int),
			}
			nodes = append(nodes, node.Right)
		}
		i++
	}
	return root
}

func Test_rob_1(t *testing.T) {
	root := NewTree([]interface{}{3, 2, 3, nil, 3, nil, 1})
	expected := 7
	ans := rob(root)
	if ans != expected {
		t.Errorf("expected: %v, rob: %v", expected, ans)
	}
}

func Test_rob_2(t *testing.T) {
	root := NewTree([]interface{}{3, 4, 5, 1, 3, nil, 1})
	expected := 9
	ans := rob(root)
	if ans != expected {
		t.Errorf("expected: %v, rob: %v", expected, ans)
	}
}
