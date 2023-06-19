/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfSubtree(root *TreeNode) int {
    _, res, _ := find(root)
    return res
}
func find(root *TreeNode) (int, int, int) {
    if root == nil {
        return 0, 0, 0
    }
    count1, res1, val1 := find(root.Left)
    count2, res2, val2 := find(root.Right)
    res := res1 + res2
    if (val1 + val2 + root.Val) / (count1 + count2 + 1)  == root.Val {

        res ++
    }
    
    return count1 + count2 + 1, res, val1 + val2 + root.Val
}