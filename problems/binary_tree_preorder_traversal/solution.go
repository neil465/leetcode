/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
    s := []*TreeNode{}
    res := []int{}
    s = append(s,root)
    for len(s) > 0 {
        n := s[len(s)-1]
        s = s[:len(s)-1]
        if n == nil{
            break
        }
        res = append(res,n.Val)
        if n.Right != nil{
            s  = append(s,n.Right)
        }
        if n.Left != nil{
            s  = append(s,n.Left)
        }
        
    }
    return res
}