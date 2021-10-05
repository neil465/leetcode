/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    m := make(map[int]*ListNode)
    curr , count := head,0
    
    for curr != nil{
        m[count] = curr
        curr,count = curr.Next , count+1 
    }
    if m[len(m)-1-n+1] == head{return head.Next}
    m[len(m)-1-n].Next = m[len(m)-1-n+2]
    return head
}