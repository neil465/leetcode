/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type F struct{
    r *TreeNode
    res string
    
}
func binaryTreePaths(rooty *TreeNode) []string {
    result := []string{}
    b := F{rooty,""}
    helper(b,&result)
    return result
}

func helper(root F , result *[]string) {
        arrow := "->"
        if len(root.res) == 0{
            arrow = ""
        }

    if root.r.Left == nil && root.r.Right == nil{
        root.res = root.res +arrow+ fmt.Sprint(root.r.Val) 
    
        *result = append(*result,root.res)
    }
    if root.r.Right != nil{
        x := root.res + arrow + fmt.Sprint(root.r.Val)  
        helper(F{root.r.Right ,x } , result)
    }
    if root.r.Left != nil{
        x := root.res + arrow + fmt.Sprint(root.r.Val) 
        helper(F{root.r.Left ,x } , result)
    }
    
}