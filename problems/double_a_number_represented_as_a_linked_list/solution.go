/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func doubleIt(head *ListNode) *ListNode {
    cur := head
    a := []*ListNode{}
    for cur != nil {
        a = append(a, cur)
        cur = cur.Next
    }
    n := 0
    for i := len(a) -1; i >= 0 ; i-- {
        a[i].Val, n = (a[i].Val * 2 + n ) % 10, (a[i].Val * 2 + n ) / 10

    }
    if n > 0 {
        v := &ListNode{n, head}
        return v
    }
    return head

}