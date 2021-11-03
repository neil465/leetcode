/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
    small := &ListNode{}
    great := &ListNode{}
    cur1 := small
    cur2 := great
    for head != nil {
        next := head.Next
        head.Next = nil
        if head.Val < x {
            cur1.Next = head
            cur1 = cur1.Next
        } else {
            cur2.Next = head
            cur2 = cur2.Next
        }
        head = next
    }
    cur1.Next = great.Next
    return small.Next
}