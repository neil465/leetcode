/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
    if root ==nil{
        return nil
    }
    queue := []*TreeNode{root}    
    res := [][]int{}
    helper(root,&res,queue)
    return res
}

func helper(root *TreeNode,result *[][]int,queue []*TreeNode){
    if len(queue) == 0{return}
    tempqueue := []*TreeNode{}
    temp := []int{}
    
    for len(queue) > 0 {
        pop := queue[0]
        queue = queue[1:]
        
        if pop.Left != nil{
            tempqueue = append(tempqueue,pop.Left)
        }
        if pop.Right != nil{
            tempqueue = append(tempqueue,pop.Right)
        }  
        temp = append(temp,pop.Val)
    }
    
    *result = append(*result,temp)
    queue = tempqueue
    
    helper(root,result,queue)
}