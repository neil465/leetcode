/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    add := 0 
    list := &ListNode{}
    c := list
    for l1 != nil || l2 != nil {
        a := 0
        if l1 != nil {
            a = l1.Val
            l1 = l1.Next
        }
        b := 0 
        if l2 != nil {
            b = l2.Val
            l2 = l2.Next
        }
        sum := add+a+b
        add = sum/ 10
        c.Next = &ListNode{Val:sum%10}
        c = c.Next
    }
    
    if add != 0 {
        c.Next = &ListNode{Val:add}
    }
   
    return list.Next
}