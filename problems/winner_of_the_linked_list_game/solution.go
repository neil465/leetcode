/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func gameResult(head *ListNode) string {
    cur := head
    a, b := 0, 0
    for cur != nil {
        if cur.Val < cur.Next.Val {
            b ++
        } else {
            a ++
        }
        cur = cur.Next.Next
    }

    if a > b {
        return "Even"
    } else if b > a {
        return "Odd"
    } 
        return "Tie"
    
}