/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
    sum := 0
    helper(root,root,&sum)
    return sum
}
func helper (root *TreeNode , parent *TreeNode, sum *int){
    if root == nil{return}
    if parent.Left == root && root.Left == nil && root.Right == nil{
        *sum += root.Val
    }
    parent = root
    helper(root.Left,parent,sum)
    helper(root.Right,parent,sum)
}