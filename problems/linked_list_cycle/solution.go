/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    if head== nil || head.Next== nil{ 
        return false
    }
    curr := head.Next
    curr2 :=  head.Next.Next
    for curr != nil && curr2 != nil{
        if curr == curr2{
            return true
        }
        if curr2.Next== nil{ 
            break
        }
        curr = curr.Next
        curr2 = curr2.Next.Next
    }
    return false
}