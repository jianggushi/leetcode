package solution

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 对于一个节点存在偷或者不偷两种情况
// 如果偷了这个节点, 那么不能偷两个子节点
// 如果不偷这个节点, 那么可以偷两个子节点
func rob(root *TreeNode) int {
	return max(robp(root)...)
}

func robp(root *TreeNode) []int {
	if root == nil {
		return []int{0, 0}
	}
	l := robp(root.Left)
	r := robp(root.Right)
	return []int{max(l...) + max(r...), root.Val + l[0] + r[0]}
}

// 存在重复计算的问题
func rob1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// r1 偷了当前节点, 左右两个子节点不能偷, 只能偷下面的
	r1 := root.Val
	if root.Left != nil {
		r1 += rob1(root.Left.Left) + rob1(root.Left.Right)
	}
	if root.Right != nil {
		r1 += rob1(root.Right.Left) + rob1(root.Right.Right)
	}
	// r2 偷两个子节点
	r2 := rob1(root.Left) + rob1(root.Right)
	return max(r1, r2)
}

// 支持可变参数的 max
func max(nums ...int) int {
	n := len(nums)
	if n == 0 {
		panic("len(nums) equal 0")
	}
	m := nums[0]
	for i := 1; i < n; i++ {
		if nums[i] > m {
			m = nums[i]
		}
	}
	return m
}

// 普通的两个参数的 max
func max1(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
