/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
    return help(root, root.Val) 
}
func help(root *TreeNode, m int) int {
    if root == nil {
        return 0
    }
    res := help(root.Left, max(root.Val, m)) + help(root.Right, max(root.Val, m))
    if root.Val >= m {
        res ++
    }
    return res
}

func max(i, j int) int {
    if i > j { return i}
    return j
}