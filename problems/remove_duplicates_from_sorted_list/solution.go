/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil || head.Next == nil{
        return head
    }
    flag := true
    for flag != false{
        f := true
        next := head.Next
        curr := head
        if head == nil || head.Next == nil{
            return head
        }
        for next != nil{

            if curr.Val == next.Val{
                f = false
                curr.Next = next.Next
                next = next.Next

            }
            curr = curr.Next
            
            if next != nil{
                next = next.Next
            }else{
                break
            }

        }
        if f{
            flag = false
        }
    }
   
    return head
}