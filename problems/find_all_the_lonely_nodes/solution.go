/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// func getLonelyNodes(root *TreeNode) []int {
//     g := []int{}
//     helper(root,&g)
//     return g
// }

// func helper(root *TreeNode , g*[]int){
//     if root == nil{return}
//     if root.Left == nil && root.Right != nil{
//         *g = append(*g,root.Right.Val)

//     }
//     if root.Left != nil && root.Right == nil{
//         *g = append(*g,root.Left.Val)

//     }
//     helper(root.Left,g)
//     helper(root.Right,g)
 
// }

func getLonelyNodes(root *TreeNode) []int {
    q,result := []*TreeNode{root},[]int{}
    for len(q) > 0{
        pop := q[0]
        q = q[1:]
        if pop.Left != nil{
            q = append(q,pop.Left)
            if pop.Right == nil{
                result = append(result , pop.Left.Val)
            }
        }
        if pop.Right != nil{
            q = append(q,pop.Right)
            if pop.Left == nil{
                result = append(result , pop.Right.Val)
            }
        }
    }
    return result
}