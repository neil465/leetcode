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
    q,level := []*TreeNode{root},0
    for len(q) > 0{
        newQ := []*TreeNode{}
        for _,i := range q{
            if i.Left == nil && i.Right == nil {
               
                return level+1
            }
            if i.Left != nil{newQ = append(newQ,i.Left)}
            if i.Right != nil{newQ = append(newQ,i.Right)}
            
        }
        if len(newQ) > 0{level++}
        q = newQ
    }
    return level
}
