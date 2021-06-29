/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicatesUnsorted(head *ListNode) *ListNode {
    m := make(map[int]int)
    curr := head
    for curr != nil{
        m[curr.Val]++
        curr = curr.Next
    }
    var prev *ListNode
    curr = head
    prev = head
    for curr != nil{
        if m[curr.Val] != 1{
            if curr == head{
                head = head.Next
                curr = head
                prev = head
                continue
            }else if curr.Next == nil{
                prev.Next = nil
                return head
            }else{
               prev.Next = curr.Next
                curr = prev.Next
                continue 
            }
            
        }
        prev = curr
        curr = curr.Next
        
    }
    return head
    
}