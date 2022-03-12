/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    
    m := map[*Node]*Node{}
    node := head
    
    for node != nil {
        m[node], node = &Node{ Val : node.Val }, node.Next
         
    }
    
    node = head
    
    for node != nil {
        m[node].Next, m[node].Random, node = m[node.Next], m[node.Random], node.Next
    }
    
    return m[head]
}