/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
    res := []float64{}
    q  := []*TreeNode{root}
    helper(root,&res,q)
    return res
}
func helper (root *TreeNode, g *[]float64 , q[]*TreeNode){
    sum := 0
    arr := []*TreeNode{}
    b := q
    for len(q)>0{
        pop := q[0]
        if pop.Left != nil{
            arr = append(arr,pop.Left)
        }
        if pop.Right != nil{
            arr = append(arr,pop.Right)
        }
        sum += pop.Val
        q =q[1:]
    }
    fmt.Println(len(q))
    *g = append(*g,float64(sum)/float64(len(b)))
    q = arr
    if len(q) > 0 {
        helper(root,g,q)   
    }
}