/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deepestLeavesSum(root *TreeNode) int {
    q := []*TreeNode{root}
    for len(q) > 0{
        flag := true
        arr := []*TreeNode{}
        sum := 0
        for _,i := range q{
            if i.Left != nil || i.Right != nil{
                flag = false
                if i.Left != nil{
                    arr = append(arr,i.Left)
                }
                if i.Right != nil{
                    arr = append(arr,i.Right)
                }
            }
            sum+= i.Val
        }
        if flag {
            return sum
        }
        q = arr
        
    }
    return -1
}