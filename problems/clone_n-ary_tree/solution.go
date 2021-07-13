/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func cloneTree(root *Node) *Node {
    if root == nil{
        return nil
    }
    copied := Node{Val:root.Val}
    helper(root,&copied)
    return &copied
}

// func helper(root *Node, copied *Node){
//     copied.Val=1
//     copied.Children = make([]*Node,len(root.Children))
//     copied.Children[0] = &Node{Val:3}
//     copied.Children[1] = &Node{Val:2}
//     copied.Children[2] = &Node{Val:4}    
// }
func helper(root *Node, copied *Node){
    //copied = &Node{Val:root.Val}
           
    if len(root.Children) == 0{
        return
    }
    
    for _,child := range root.Children{
        copied.Children = append(copied.Children,&Node{Val:child.Val})    
    }
    
    for i,child := range root.Children{
        helper(child,copied.Children[i])
    }
}