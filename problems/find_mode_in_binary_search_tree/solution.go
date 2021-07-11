/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findMode(root *TreeNode) []int {
    g:= make(map[int]int)
    helper(root,&g)
    max := 0
    maxarr := []int{}
    
    for k,v := range g{
        if v == max{
            maxarr = append(maxarr,k)
        }
        if v > max{
            maxarr = []int{k}
            max = v
        }
    }
    return maxarr
}
func helper (root *TreeNode,g *map[int]int){
    if root == nil{return}
    
    x := *g
    x[root.Val]++
    helper(root.Left,g)
    helper(root.Right,g)
}