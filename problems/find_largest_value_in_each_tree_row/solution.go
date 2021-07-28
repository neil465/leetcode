/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestValues(root *TreeNode) []int {
    if root == nil{return []int{}}
    res := []int{}
    q := []*TreeNode{root}
    
    for len(q) > 0{
        max := -10000000000000
        arr := []*TreeNode{}
        for _,i := range q{
            if i.Val > max{
                max = i.Val
            }
            if i.Left != nil{
                arr = append(arr,i.Left)
            }
            if i.Right != nil{
                arr = append(arr,i.Right)
            }
        }
        q = arr
        res = append(res,max)
    }
    return res
}