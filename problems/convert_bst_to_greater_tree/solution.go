/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func convertBST(root *TreeNode) *TreeNode {
    g := []int{}
    helper(root,&g)
    sort.Ints(g)
    m := make(map[int]int)
    for i := 0 ; i<len(g) ; i++{
        sum := 0
        for _,i := range g[i:]{
            sum += i
        }
        m[g[i]] = sum
    }
    helper2(root,m)
    return root
}
func helper(root *TreeNode , g *[]int){
    if root == nil{return}
    *g = append(*g,root.Val)
    helper(root.Left,g)
    helper(root.Right,g)
}
func helper2(root *TreeNode , m map[int]int){
    if root == nil{return}
    root.Val = m[root.Val]
    helper2(root.Left,m)
    helper2(root.Right,m)
}