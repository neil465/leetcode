func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	stack := []*TreeNode{root1}
	arr := []int{}
	arr2 := []int{}
	for len(stack) > 0 {
		value := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if value != nil {
			if value.Left == nil && value.Right == nil {
				arr = append(arr, value.Val)
			} else {
				stack = append(stack, value.Left)
				stack = append(stack, value.Right)
			}
		}
	}
	stack = []*TreeNode{root2}
	for len(stack) > 0 {
		value := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if value != nil {
			if value.Left == nil && value.Right == nil {
				arr2 = append(arr2, value.Val)
			} else {
				stack = append(stack, value.Left)
				stack = append(stack, value.Right)
			}
		}
	}
    fmt.Println(arr,arr2)
    
    if len(arr) != len(arr2) {
        return false
    }
    
	for i := 0; i < len(arr); i++ {
		if arr[i] != arr2[i] {
			return false
		}
	}
	return true
}
