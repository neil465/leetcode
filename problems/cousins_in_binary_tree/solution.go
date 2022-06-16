/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isCousins(root *TreeNode, x int, y int) bool {
    p1, l1 := findParentAndLevel(nil, root, x, 1)
    p2, l2 := findParentAndLevel(nil, root, y, 1)
    return p1 != p2 && l1 == l2
}
func findParentAndLevel(parent, root *TreeNode, val, level int) (*TreeNode, int) {
    if root == nil {
        return nil, -1
    }
    if root.Val == val {
        return parent,level
    }
    if p, l := findParentAndLevel(root, root.Right, val, level + 1); p != nil {
        return p, l
    }
    if p, l := findParentAndLevel(root, root.Left, val, level + 1); p != nil {
        return p, l
    }
    return nil, -1
}