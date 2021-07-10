/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func searchBST(root *TreeNode, val int) *TreeNode {
    if root == nil{
        return nil
    }
    if root.Val == val {
        return root
    }
    l1 := searchBST(root.Left,val)
    if l1 != nil{
        return l1
    }
    l2 := searchBST(root.Right,val)
    if l2 != nil{
        return l2
    }
    return nil
}
