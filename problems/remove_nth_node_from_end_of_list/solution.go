/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    if n == 0 || head == nil{
        return head
    }
    if head.Next == nil{
        return nil
    }
    
    count := 1
    c := head
    val := 1
    
    for c != nil{
        c = c.Next
        val++
    }
    fmt.Println(val)
    if n == val-1{
        head = head.Next
        return head
    }
    val -= n+1
    
    cur := head
    for cur != nil{
        if count == val{
            cur.Next = cur.Next.Next
            return head
        }
        count ++
        cur = cur.Next
    }
    return head
    
}