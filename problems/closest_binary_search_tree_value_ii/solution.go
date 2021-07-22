/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func closestKValues(root *TreeNode, target float64, k int) []int {
    g := []int{}
    helper(root,&g)
    sort.Slice(g,func(i,j int)bool{
        if abs(target-float64(g[i])) < abs(target-float64(g[j])){
            return false
        }
        return true
    })
    return g[len(g)-k:]
}
func helper(root *TreeNode , g *[]int){
    if root == nil{return}
    *g = append(*g,root.Val)
    helper(root.Left,g)
    helper(root.Right,g)
}
func abs(i float64)float64{
    if i <0 {
        return i*-1
    }
    return i
}