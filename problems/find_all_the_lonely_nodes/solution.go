/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getLonelyNodes(root *TreeNode) []int {
    g := []int{}
    helper(root,&g)
    return g
}

func helper(root *TreeNode , g*[]int){
    if root == nil{return}
    if root.Left == nil && root.Right != nil{
        *g = append(*g,root.Right.Val)

    }
    if root.Left != nil && root.Right == nil{
        *g = append(*g,root.Left.Val)

    }
    helper(root.Left,g)
    helper(root.Right,g)
 
}