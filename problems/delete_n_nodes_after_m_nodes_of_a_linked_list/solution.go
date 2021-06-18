/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNodes(head *ListNode, m int, n int) *ListNode {
    curr := head
    
    var l *ListNode
    var curr2 *ListNode
    count := 0
    flag := false
    for curr != nil{
        if  flag == true{
            count++
            if count==n{
                flag = false 
                count = 0
            } 
        }else{
            if l == nil {
                l = &ListNode{Val:curr.Val}
                curr2 = l
            }else{
                l1 := &ListNode{}
                l1.Val = curr.Val
                curr2.Next = l1
                curr2 = curr2.Next
            }
            count++
            if count==m{
                flag=true
                count = 0
            } 
        }
        
        curr = curr.Next
    }
    return l
    
}