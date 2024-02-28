/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func findBottomLeftValue(root *TreeNode) int {
    q := [][]*TreeNode{{root}}
    for len(q[len(q) - 1]) > 0 {
        nextList := []*TreeNode{}

        for _, v := range q[len(q) - 1] {
            if v.Left != nil {
                nextList = append(nextList, v.Left)
            }
            if v.Right != nil {
                nextList = append(nextList,  v.Right)
            }
        }
        
        q = append(q, nextList)

    }
    return q[len(q) - 2][0].Val
}