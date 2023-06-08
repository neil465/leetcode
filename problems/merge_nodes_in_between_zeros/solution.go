/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeNodes(head *ListNode) *ListNode {
    cur := head.Next
    a := &ListNode{Val:0}
    acur := a
    s := 0
    for cur != nil {
        if cur.Val == 0 {
            acur.Next = &ListNode{Val:s} 
            s = 0
            acur = acur.Next
        } else {
            s += cur.Val
        }
        cur = cur.Next
    }
    return a.Next
}