/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    t,h := head,head
    for h != nil{
        h = h.Next
        if h != nil{
            h = h.Next
        }else{
            break
        }
        t = t.Next
    }
    return t
}