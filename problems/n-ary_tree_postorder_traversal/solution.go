/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func postorder(root *Node) []int {
    g := []int{}
    helper(root,&g)
    return g
}
func helper(root *Node , g *[]int){
    if root == nil{return}
    
    for _,i := range root.Children{
        helper(i,g)
    }
    *g = append(*g,root.Val)
}