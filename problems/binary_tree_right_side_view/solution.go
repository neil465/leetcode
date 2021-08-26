/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
    if root == nil{return nil}
    res := []int{}
    q := []*TreeNode{root}
    for len(q) > 0{
        temp := []*TreeNode{}
        res = append(res,q[len(q)-1].Val)
        for _,i := range q{
            if i.Left != nil{temp = append(temp,i.Left)}
            if i.Right != nil{temp = append(temp,i.Right)}            
        }
        q = temp
    }
    return res
}