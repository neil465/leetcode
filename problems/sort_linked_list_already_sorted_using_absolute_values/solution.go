/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortLinkedList(head *ListNode) *ListNode {
    prev,cur := head , head.Next 
    for cur != nil{
        if cur.Val <0 {
            prev.Next = cur.Next
            cur.Next = head
            head = cur
            cur = prev.Next
        }else{
            prev = cur
            cur = cur.Next
        }
    }
    return head
}