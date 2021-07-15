/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    if root == nil{return 0}
    q := []*TreeNode{root}
    level := 1
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
        }
        if pop.Right != nil{
            arr= append(arr,pop.Right)
        }
        
    }
    if len(arr)> 0{
        *level++
        helper(arr,level)
    }
    
}