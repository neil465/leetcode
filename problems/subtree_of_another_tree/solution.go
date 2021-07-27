// /**
//  * Definition for a binary tree node.
//  * type TreeNode struct {
//  *     Val int
//  *     Left *TreeNode
//  *     Right *TreeNode
//  * }
//  */
// func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
//     q := []*TreeNode{root}
    
//     for len(q) > 0 {
//         pop := q[0]
//         q = q[1:]
//         if pop.Val == subRoot.Val{
//             g,g2 := []int{},[]int{}
//             helper(pop,&g)
//             helper(subRoot,&g2)
//             if len(g) == len(g2){
//                 fmt.Println(g,g2)
                
//                 if reflect.DeepEqual(g,g2){
//                     return true
//                 }
//             }
//         }
//         if pop.Left != nil{
//             q = append(q,pop.Left)
//         }
//         if pop.Right != nil{
//             q = append(q,pop.Right)
//         }
//     }
//     return false
// }
// func helper(root *TreeNode , g *[]int){
//     if root.Left != nil{
//         helper(root.Left,g)
//         *g = append(*g,root.Val)
//     }else{
//         *g = append(*g,0)
//     }
//     if root.Right != nil{
//         helper(root.Right,g)
//         *g = append(*g,root.Val)
//     }else{
//         *g = append(*g,0)
//     }
    
// }
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    q := []*TreeNode{root}
    
    for len(q) > 0 {
        pop := q[0]
        q = q[1:]
        if pop.Val == subRoot.Val{
            g,g2 := []int{},[]int{}
            p,s := []*TreeNode{pop},[]*TreeNode{subRoot}
            helper(p,&g)
            helper(s,&g2)
            if reflect.DeepEqual(g,g2){
                return true
            }
        }
        if pop.Left != nil{
            q = append(q,pop.Left)
        }
        if pop.Right != nil{
            q = append(q,pop.Right)
        }
    }
    return false
}
func helper(q []*TreeNode , g *[]int){
    if len(q)  == 0{return}
    pop := q[0]
    q = q[1:]
    if pop.Left != nil{
        *g = append(*g,pop.Left.Val)
        q = append(q,pop.Left)
    }
    if pop.Left == nil{
        *g = append(*g,-100002)
    }
    if pop.Right != nil{
        *g = append(*g,pop.Right.Val)
        q = append(q,pop.Right)
    }
    if pop.Right == nil{
        *g = append(*g,-100002)
    }
    helper(q,g)
}