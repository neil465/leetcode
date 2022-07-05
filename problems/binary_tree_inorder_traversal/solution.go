/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
    if root.Left == nil && root.Right == nil {
        return []int{root.Val}
    }
    return append(append(inorderTraversal(root.Left), root.Val), inorderTraversal(root.Right)...)
}