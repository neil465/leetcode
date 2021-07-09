/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
    g := []int{}
    iteration(root,&g)
    return g
}

func iteration(root *TreeNode, g *[]int){
    if root == nil{
        return
    }
    *g = append(*g,root.Val)
    iteration(root.Left,g)
    iteration(root.Right,g)
    
}