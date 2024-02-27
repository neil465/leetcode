/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var best = 0
func diameterOfBinaryTree(root *TreeNode) int {
    best = 0
    re(root)
    return best
}
func re(cur *TreeNode) int {
    if cur == nil { return 0 }
    v1, v2 := re(cur.Left), re(cur.Right)
    best = max(v1 + v2 , best)
    return max(v1, v2) + 1
}