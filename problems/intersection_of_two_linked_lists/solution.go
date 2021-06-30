/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    m := make(map[*ListNode]bool)
    c := headA
    c2 := headB
    for c != nil || c2 != nil{
        if c == nil{
            if m[c2] == true{
                return c2
            }
            m[c2] = true
            c2 = c2.Next
            continue
            
        }
        if c2 == nil{
            if m[c] == true{
                return c
            }
            m[c] = true
            c = c.Next
            continue
        }
        if m[c] == true{
            return c
        }
        if m[c2] == true{
            return c2
        }
        if c2 == c{
            return c2
        }
        fmt.Println(c.Val,c2.Val)
        m[c]= true
        m[c2] = true
        c = c.Next
        c2 = c2.Next
    }
    return nil
}