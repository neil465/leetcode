/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func closestValue(root *TreeNode, target float64) int {
    g := 10000000000
    helper(root,target,&g)
    return g
}

func helper(root *TreeNode,target float64,g *int){
    if root == nil{
        return
    }
    if abs(float64(root.Val)-target) < abs(float64(*g)-target){
        
        *g = root.Val
    }
    helper(root.Left,target,g)
    fmt.Println(*g)
    helper(root.Right,target,g)
    fmt.Println(*g)
}

func abs(i float64) float64{
    if i <0 {
        return i*-1
    }
    return i
}