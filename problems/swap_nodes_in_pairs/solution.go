/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    curr := head
    count := 0
    for curr.Next != nil{
        if count %2 == 0{
            curr.Val,curr.Next.Val = curr.Next.Val,curr.Val
        }
        curr = curr.Next
        count++
    }
    return head
}