/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type F struct{
    r *TreeNode
    res []int
    sum int
}
func pathSum(root *TreeNode, targetSum int) [][]int {
    if root == nil{
        return [][]int{}
    }
    res := []int{}
    sum := 0
    f := F{root,res,sum}
    result := [][]int{}
    helper(f,&result,targetSum)
    
    return result
}

func helper(root F , result *[][]int,target int) {
    root.sum += root.r.Val
    temp := []int{}
    temp = append(temp,root.res...)
    temp = append(temp,root.r.Val)
    root.res = temp
    if root.r.Left == nil && root.r.Right == nil && root.sum == target{
        *result = append(*result,root.res)
    }
    
    if root.r.Right != nil{
         
        helper(F{root.r.Right ,root.res,root.sum } , result,target)
    }
    if root.r.Left != nil{
    
         helper(F{root.r.Left ,root.res,root.sum },  result,target)
    }
}