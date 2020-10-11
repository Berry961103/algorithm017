package week02

// 前序遍历 递归
func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return nil
	}
	res = append(res, root.Val)
	res = append(res, preorderTraversal(root.Left)...)
	res = append(res, preorderTraversal(root.Right)...)
	return res
}

// 前序遍历 栈 迭代
func preorderTraversalStack(root *TreeNode) (res []int) {
	stack := []*TreeNode{}
	for len(stack) > 0 || root != nil {
		for root != nil {
			res = append(res, root.Val)
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = root.Right

	}
	return res
}
