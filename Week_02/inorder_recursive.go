package week02

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 中序遍历 递归
func inorderRecursive(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	// rest := append(inorderRecursive(root.Left), root.Val)
	// rest = append(rest, inorderRecursive(root.Right)...)
	res := inorderRecursive(root.Left)
	res = append(res, root.Val)
	resR := inorderRecursive(root.Right)
	res = append(res, resR...)
	return res
}

// 中序遍历 栈 迭代
func inorderTraversaStack(root *TreeNode) (res []int) {
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		res = append(res, root.Val)
		stack = stack[:len(stack)-1]
		root = root.Right
	}
	return
}
