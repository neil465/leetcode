/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil && l2 == nil{
        return nil
    }
    cy,cx :=  l2,l1
    l := &ListNode{}
    c := l
    var prev *ListNode
    for cy != nil || cx != nil{
        
        if cy == nil && cx != nil{
            c.Val = cx.Val
		    l1 := &ListNode{}
            prev = c
		    c.Next =l1
		    c = c.Next
            cx = cx.Next
        }else if cx == nil && cy != nil{
            c.Val = cy.Val
		    l1 := &ListNode{}
            prev = c
		    c.Next =l1
		    c = c.Next
            cy = cy.Next
        }else if  cy.Val > cx.Val{
            
            c.Val = cx.Val
		    l1 := &ListNode{}
            prev = c
		    c.Next =l1
		    c = c.Next
            cx = cx.Next
            
        }else if cy.Val == cx.Val {
            c.Val = cx.Val
		    l1 = &ListNode{}
            prev = c
		    c.Next =l1
		    c = c.Next
            cx = cx.Next
            c.Val = cy.Val
		    l1 = &ListNode{}
            prev = c
		    c.Next =l1
		    c = c.Next
            cy = cy.Next
        }else{
            c.Val = cy.Val
		    l1 := &ListNode{}
		    c.Next =l1
            prev = c
		    c = c.Next
            cy = cy.Next
            
            
        }
        
        
        
        
        
    }
    prev.Next = nil
    return l
}