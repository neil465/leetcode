/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
     g := []int{}
    iteration(root,&g)
    return g
}

func iteration(root *TreeNode,g *[]int) {
    if root == nil{
        return
    }
   
    iteration(root.Left,g)
    iteration(root.Right,g)
     *g = append(*g,root.Val)

}