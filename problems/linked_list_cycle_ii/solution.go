/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
    slow,fast := head,head
    visited := map[*ListNode]bool{}
    for fast != nil{
        
        
        visited[slow] = true
        slow = slow.Next
        fast = fast.Next
        if visited[fast] {
            return fast
        }
        if fast != nil{
            fast = fast.Next
        }
        if visited[fast] {
            return fast
        }
        
    }
    return nil
}