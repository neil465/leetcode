/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var res = 0

func maxAncestorDiff(root *TreeNode) int {
    res = 0
    a(root)
    return res
}
func a(root *TreeNode) (int, int) {
    if root == nil {
        return math.MaxInt32, math.MinInt32
    }
    tminLeft, tmaxLeft := a(root.Left)
    tminRight, tmaxRight := a(root.Right)
    bmin, bmax := min(root.Val, tminLeft, tminRight), max(root.Val, tmaxLeft, tmaxRight)

    res = max(res, abs(bmin - root.Val), abs(bmax - root.Val))
    return bmin, bmax
}
func abs(i int) int {
    if i > 0 {
        return i
    }
    return -i
}