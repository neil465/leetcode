/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
    if head == nil{
        return nil
    }
    k := &ListNode{}
	c := k
	var prev *ListNode
    curr := head
    for curr != nil{
        if curr.Val != val{
            c.Val = curr.Val
		    l1 := &ListNode{}
		    prev = c
		    c.Next =l1
		    c = c.Next
        }
        curr = curr.Next
    }
    if k.Next!= nil{
        fmt.Println(k)
        prev.Next = nil 
        return k
    }
    return nil
}