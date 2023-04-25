/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumRootToLeaf(root *TreeNode) int {
    return sum(root, fmt.Sprint(root.Val))
}
func sum(root *TreeNode, curbn string) int {
    if root.Left == nil && root.Right == nil {

        i, _ := strconv.ParseInt(curbn , 2, 64)

        return int(i)
    }
    s := 0
    if root.Left != nil {
        s = sum(root.Left, curbn + fmt.Sprint(root.Left.Val)) + s
    }
    if root.Right != nil {
        s = sum(root.Right, curbn + fmt.Sprint(root.Right.Val)) + s
    }
    return s
}