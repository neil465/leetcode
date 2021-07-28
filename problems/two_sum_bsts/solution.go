/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func twoSumBSTs(root1 *TreeNode, root2 *TreeNode, target int) bool {
    flag,m := false,make(map[int]bool)
    helper(root1,m)
    helper2(root2,m,&flag,target)
    return flag
}
func helper(root *TreeNode,m map[int]bool){
    if root == nil{return}
    m[root.Val] = true
    helper(root.Left,m)
    helper(root.Right,m)
}
func helper2(root *TreeNode,m map[int]bool,flag *bool,target int){
    if root == nil{return}
    
    if m[target - root.Val] == true{
        
        *flag = true
    }
    helper2(root.Left,m,flag,target)
    helper2(root.Right,m,flag,target)
}