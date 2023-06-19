/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func tree2str(root *TreeNode) string {
    if root == nil {
        return ""
    }
    a, b, res :=  tree2str(root.Left),  tree2str(root.Right), fmt.Sprint(root.Val) 
    if a != "" && b == "" {
        return res + "(" + a + ")"
    } else if a == "" && b == "" { 
        return res 
    }
    return res + "(" + a + ")" + "(" + b + ")"
}