/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
    if root == nil{return 0}
    level := 1
    q := []*TreeNode{root}
    helper(q,&level)
    return level
}

func helper (q []*TreeNode , level *int){
    arr := []*TreeNode{}
    for len(q) > 0 {
        pop := q[0]
        q = q[1:]
        if pop.Left != nil{
            arr= append(arr,pop.Left)
            
            if pop.Left.Right == nil && pop.Left.Left == nil{
                *level++
                return
            }
        }
        if pop.Right != nil{
            arr= append(arr,pop.Right)
            if pop.Right.Left == nil && pop.Right.Right == nil{
                *level++
                return
            }
        }
        
    }
    if len(arr)> 0{
        *level++
        helper(arr,level)
    }
    
}