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
    res int
}
func hasPathSum(root *TreeNode, targetSum int) bool {
    if root == nil{
        return false
    }
    f := F{r:root}
    flag := false
    helper(f,targetSum,&flag)
    return flag
}
func helper(f F , targetSum int, flag *bool){
    
    f.res += f.r.Val
    if f.r.Left == nil && f.r.Right == nil && f.res == targetSum{
        *flag = true
        return 
    }
    
    if f.r.Left != nil{
        helper(F{r:f.r.Left,res:f.res},targetSum,flag)
    }
    
    if f.r.Right != nil{       
        helper(F{r:f.r.Right,res:f.res},targetSum,flag)
    }
}
