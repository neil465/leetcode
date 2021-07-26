/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// func maxDepth(root *TreeNode) int {
//     if root == nil{
//         return 0
//     }
//     level := 1
//     q := []*TreeNode{root}
//     for len(q) > 0{
//         arr := []*TreeNode{}
//         for _,i := range q{
//             if i.Left != nil {arr = append(arr,i.Left)}
//             if i.Right != nil {arr = append(arr,i.Right)}
//         }
//         if len(arr) > 0{level++}
//         q = arr
//     }
//     return level

// }
func maxDepth(root *TreeNode) int {
    if root == nil {return 0}
    return 1+max(maxDepth(root.Left),maxDepth(root.Right))
}
func max (i,j int)  int{
    if i > j{
        return i 
    }
    return j
}