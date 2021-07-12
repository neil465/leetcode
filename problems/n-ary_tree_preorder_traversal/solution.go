/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
    g := []int{}
    helper(root,&g)
    return g
}
func helper(root *Node,g *[]int){
    if root == nil{return}
    *g = append(*g,root.Val)
    for _,i := range root.Children{
        helper(i,g)
    }
}