/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
    if root != nil { return 1 + countNodes(root.Right) + countNodes(root.Left) }else{return 0}
}