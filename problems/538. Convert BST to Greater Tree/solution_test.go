package solution

import (
	"fmt"
	"testing"
)

func (t *TreeNode) String() string {
	ans := []interface{}{}
	if t == nil {
		return fmt.Sprintf("%v", ans)
	}
	nodes := []*TreeNode{}
	nodes = append(nodes, t)
	for len(nodes) != 0 {
		node := nodes[0]
		nodes = nodes[1:]
		if node != nil {
			ans = append(ans, node.Val)
			nodes = append(nodes, node.Left, node.Right)
		} else {
			ans = append(ans, nil)
		}
	}
	return fmt.Sprintf("%v", ans)
}

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

func EqualTree(t1 *TreeNode, t2 *TreeNode) bool {
	eq := false
	preOrder(t1, t2, &eq)
	return eq
}

func preOrder(t1 *TreeNode, t2 *TreeNode, eq *bool) {
	if t1 == nil && t2 == nil {
		*eq = true
		return
	}
	if t1 == nil || t2 == nil {
		*eq = false
		return
	}
	if t1.Val != t2.Val {
		*eq = false
		return
	}
	preOrder(t1.Left, t2.Left, eq)
	preOrder(t1.Left, t2.Left, eq)
}

func Test_convertBST1(t *testing.T) {
	root := NewTree([]interface{}{5, 2, 13})
	expected := NewTree([]interface{}{18, 20, 13})
	ans := convertBST(root)
	if !EqualTree(ans, expected) {
		t.Errorf("expected: %v, convertBST: %v", expected, ans)
	}
}
