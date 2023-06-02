/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
    return findValid(root, -123, 123)
}
func findValid(root *TreeNode, mi, ma int) bool {
    if root == nil { return true }
    if ( mi != -123 && root.Val <= mi ) || (ma != 123 &&  root.Val >= ma) { return false }
    return findValid(root.Left, mi, root.Val) && findValid(root.Right, root.Val, ma)
}