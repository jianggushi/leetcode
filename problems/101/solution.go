package solution

//  Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 迭代解法
// 按层遍历，模拟两棵树
// 一棵树先左后右遍历，一棵树先右后左遍历
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	q1 := []*TreeNode{root}
	q2 := []*TreeNode{root}

	for len(q1) != 0 && len(q2) != 0 {
		node1, node2 := q1[0], q2[0]
		q1, q2 = q1[1:], q2[1:]
		if node1 == nil && node2 == nil {
			continue
		}
		if node1 == nil || node2 == nil {
			return false
		}
		if node1.Val != node2.Val {
			return false
		}
		q1 = append(q1, node1.Left, node1.Right)
		q2 = append(q2, node2.Right, node2.Left)
	}
	if len(q1) != 0 || len(q2) != 0 {
		return false
	}
	return true
}

// 前面的解法对 root 节点做了特殊处理
// 其实也可以统一掉，想像 root 上面还有一个节点
// 把当前的树复制一份当做更上面一个节点的两个子树
func isSymmetric2(root *TreeNode) bool {
	return symmetric(root, root)
}

// 递归解法
// 树本身就是一个递归结构，很容易要想用递归解法
// 要判断一棵树是不是镜像的，要看他的左子树和右子树是不是镜像的
func isSymmetric1(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return symmetric(root.Left, root.Right)
}

func symmetric(r1, r2 *TreeNode) bool {
	if r1 == nil && r2 == nil {
		return true
	}
	if r1 == nil || r2 == nil {
		return false
	}
	if r1.Val != r2.Val {
		return false
	}
	return symmetric(r1.Left, r2.Right) && symmetric(r1.Right, r2.Left)
}
