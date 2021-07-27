/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTarget(root *TreeNode, k int) bool {
    m,res := make(map[int]bool),false
    h(root,m,k,&res)
    return res
}
func h (root *TreeNode , m map[int]bool, k int,res *bool){
    if root == nil{return}
    if m[k-root.Val] == true{*res = true}
    m[root.Val] = true
    h(root.Left,m,k,res)
    h(root.Right,m,k,res)
}