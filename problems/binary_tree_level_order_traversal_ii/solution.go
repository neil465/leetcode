/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
    if root ==nil{
        return nil
    }
    q := []*TreeNode{root}    
    res := [][]int{}
    helper(root,&res,q)
    return res
}

func helper(root *TreeNode,g *[][]int,q []*TreeNode){
    arr := []*TreeNode{}
    temp := []int{}
    
    for len(q) > 0 {
        pop := q[0]
        if pop.Left != nil{
            arr = append(arr,pop.Left)
        }
        if pop.Right != nil{
            arr = append(arr,pop.Right)
        }
        
        q = q[1:]
        temp = append(temp,pop.Val)
    }
    
    *g = append([][]int{temp}, *g...)
    q = arr
    if len(q) > 0{
        helper(root,g,q)
    }
}