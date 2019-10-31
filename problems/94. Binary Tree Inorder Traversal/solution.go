package solution

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 方法二：迭代方法
// 利用栈，从根节点开始，先把左孩子全部入栈
// 然后再一次出栈，访问该节点，然后再对右节点进行上述操作
func inorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	// 用 slice 构造栈
	stack := make([]*TreeNode, 0)
	p := root
	for p != nil || len(stack) != 0 {
		for p != nil {
			// 入栈
			stack = append(stack, p)
			p = p.Left
		}
		// 出栈
		p = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, p.Val)
		p = p.Right
	}
	return ans
}

// 方法一：递归方法
func inorderTraversal1(root *TreeNode) []int {
	ans := make([]int, 0)
	inorder(root, &ans)
	return ans
}

func inorder(root *TreeNode, ans *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, ans)
	// 注意：slice []int 在 append 时可能会发生重新分配内存
	// 所以传参时要使用 *[]int
	*ans = append(*ans, root.Val)
	inorder(root.Right, ans)
}
