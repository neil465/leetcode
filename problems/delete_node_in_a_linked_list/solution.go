/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {
    cur := node
    var prev *ListNode
    for cur.Next != nil{
        cur.Val = cur.Next.Val
        prev = cur
        cur = cur.Next
        fmt.Println(cur)
    }
    prev.Next = nil
}