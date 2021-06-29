/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
    m := make(map[*ListNode]bool)
    if head == nil || head .Next == nil{
        return nil
    }
    curr := head
    for curr != nil{
        if m[curr] == true{
            return curr
        }
        m[curr] = true
        curr = curr.Next
    }
    return nil
}