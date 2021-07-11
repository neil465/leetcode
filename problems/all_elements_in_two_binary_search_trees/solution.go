/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
    g1 := []int{}
    g2 := []int{}
    helper(root1,&g1)
    helper(root2,&g2)
    g1 = append(g1,g2...)
    sort.Ints(g1)
    return g1
}
func helper(root *TreeNode,g *[]int,){
    if root == nil{return}
    *g = append(*g,root.Val)
    helper(root.Left,g)
    helper(root.Right,g)
}