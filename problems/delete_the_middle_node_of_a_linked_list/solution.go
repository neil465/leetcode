/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteMiddle(head *ListNode) *ListNode {
    if head == nil{return nil}
    if head.Next == nil{return nil}
    slow , fast := head,head
    var prev *ListNode
    var prprev *ListNode
    for fast != nil{
        prprev = prev
        prev = slow
        slow = slow.Next
        fast = fast.Next
        if fast != nil{
            fast = fast.Next
            if fast == nil{
                prev.Next = prev.Next.Next
                return head
            }
        }
    }
    prprev.Next = prprev.Next.Next 
    return head
}