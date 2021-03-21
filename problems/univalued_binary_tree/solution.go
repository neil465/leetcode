func isUnivalTree(root *TreeNode) bool {
	same := root
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		value := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if value != nil {
			if same.Val != value.Val {
				return false
			}
			stack = append(stack, value.Left, value.Right)
		}

	}
	return true

}