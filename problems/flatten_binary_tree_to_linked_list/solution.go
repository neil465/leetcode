/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var p *TreeNode
func flatten(root *TreeNode)  {
    p = nil
    f(root)
}
func f(root *TreeNode) {
    if root == nil {
        return
    }
    f(root.Right)
    f(root.Left)
    root.Right = p
    root.Left = nil
    p = root
}
