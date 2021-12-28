/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    slow , fast := head,head
    var prev *ListNode
    for fast != nil{
        prev = slow
        slow = slow.Next
        fast = fast.Next
        if fast != nil{
            fast = fast.Next
            if fast == nil{return slow}
        }
    }
    return prev
}