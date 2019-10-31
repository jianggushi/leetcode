package solution

// TreeNode definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 方法一：
// 以每一个节点作为根节点，前序遍历二叉树
// 在遍历过程中如果遇到使 sum 等于 0 的就算作一条 path
// 参考：https://leetcode-cn.com/problems/path-sum-iii/solution/leetcode-437-path-sum-iii-by-li-xin-lei/

func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	return path(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}

func path(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	res := 0
	if root.Val == sum {
		res++
	}
	res += path(root.Left, sum-root.Val)
	res += path(root.Right, sum-root.Val)
	return res
}
