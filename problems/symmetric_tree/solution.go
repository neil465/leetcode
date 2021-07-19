/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    flag := true
    q := []*TreeNode{root}
    helper(root,&flag,q)
    return flag
}
func helper(root *TreeNode,flag *bool,queue []*TreeNode){
    if len(queue) == 0{return}
    tempqueue := []*TreeNode{}
    temp := []*TreeNode{}
    
    for len(queue) > 0 {
        pop := queue[0]
        queue = queue[1:]
        
        if pop.Left != nil{
            tempqueue = append(tempqueue,pop.Left)
            temp = append(temp,pop.Left)
        }else{
            temp = append(temp,pop.Left)
        }
        if pop.Right != nil{
            tempqueue = append(tempqueue,pop.Right)
            temp = append(temp,pop.Right)
        }else{
            temp = append(temp,pop.Right)
        }
        
        
    }
    l,r := 0,len(temp)-1
    for l<= r{
        if temp[l] == nil && temp[r] != nil{
            *flag = false
            return
        }else if temp[r] == nil && temp[l] != nil{
            *flag = false
            return
        }else if temp[l] != nil && temp[r] != nil && temp[r].Val != temp[l].Val {
            *flag = false
            return
        }
        l++
        r--
    }
    
    helper(root,flag,tempqueue)
}