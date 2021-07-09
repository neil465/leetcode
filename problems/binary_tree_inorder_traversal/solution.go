/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    g := []int{}
    
    Inorder(root,&g)
    return g
}
func Inorder(root *TreeNode,g *[]int){
    if root == nil{return}
    Inorder(root.Left,g)
    *g = append(*g,root.Val)
    Inorder(root.Right,g)
}
