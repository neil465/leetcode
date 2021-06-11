func isUnivalTree(root *TreeNode) bool {

	same := root
	stack := []*TreeNode{root}
    for len(stack) != 0{
        n := stack[0]
        stack =stack[1:]
        
        if n.Val != same.Val && n != nil{
            
            return false
        }
        if n.Right != nil{
            stack = append(stack,n.Right)
        }
        if n.Left != nil{
            stack = append(stack,n.Left)
        }
    }
    return true

}