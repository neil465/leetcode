/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
    curr := list1
    curr2 := list2
    var endofL2 *ListNode
    count := 0
    
    for curr2 != nil{
        if curr2.Next == nil{
            endofL2 = curr2
        }
        curr2 = curr2.Next
    }
    
    for curr != nil{
        if count == a-1{
            afterRemoved := curr
            
            for i := 0 ; i < b-a+2 ; i++{
                afterRemoved = afterRemoved.Next
            }
            
            curr.Next = list2
            endofL2.Next = afterRemoved
            return list1
        }
        count++
        curr = curr.Next
    }
    return &ListNode{}
}