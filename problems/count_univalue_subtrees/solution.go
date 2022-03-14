/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func countUnivalSubtrees(root *TreeNode) int {
    result,_ := count(root)
    return result
}
func count(root *TreeNode) (int, bool) {
    if root == nil {
        return 0, true
    }
    
    countLeft, ifLeft := count(root.Left)
    countRight, ifRight := count(root.Right)
    sum := countLeft + countRight
    if (root.Left == nil ||  ((root.Left.Val == root.Val) && ifLeft )) && (root.Right == nil ||  ((root.Right.Val == root.Val) && ifRight)) {
        return 1 + sum, true
    }
    return sum, false
    
}