/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }
    q := []*TreeNode{root}
    result := [][]int{[]int{root.Val}}
    
    for len(q) > 0 {
        temp := []int{}
        n := len(q)
        for i := 0 ; i < n ; i++{
            p := q[0]
            q = q[1:]

            if p.Left != nil {
                q = append(q, p.Left)
                temp = append(temp, p.Left.Val)
            }
            if p.Right != nil {
                q = append(q, p.Right)
                temp = append(temp, p.Right.Val)
            }
        }
        if len(temp) > 0 {
            result = append(result, temp)
        }
        
    }
    return result
    
}