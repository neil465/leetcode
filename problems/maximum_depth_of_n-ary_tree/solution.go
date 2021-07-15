/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func maxDepth(root *Node) int {
    if root == nil{
        return 0
    }
    level := 1
    q := []*Node{root}
    helper(&level,q)
    return level
}
func helper(level *int , q []*Node){
    arr := []*Node{}
    
    for len(q) > 0 {
        pop := q[0]
        q  = q[1:]
        arr = append(arr,pop.Children...)
    }
    if len(arr) > 0{ 
        *level ++
        helper(level,arr)
    }
}