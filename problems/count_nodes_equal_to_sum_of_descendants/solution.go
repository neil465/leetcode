/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func equalToDescendants(root *TreeNode) int {
    a ,_ := h(root)
    return a
}
func h(root *TreeNode) (int, int) {
    if root == nil { return 0, 0 }

    c1, v1 := h(root.Left)
    c2, v2 := h(root.Right)
    if v1 + v2 == root.Val {
        c1 ++
    }
    return c1 + c2, v1 + v2 + root.Val

}