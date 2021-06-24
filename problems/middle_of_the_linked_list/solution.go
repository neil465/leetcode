/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// func middleNode(head *ListNode) *ListNode {
//     m := make(map[int]*ListNode)
//     count := 0
//     for head != nil{
//         m[count] = head
//         head = head.Next
//         count++
//     }
//     return m[count/2]
// }
func middleNode(head *ListNode) *ListNode {
    count := 0
    c := head
    for c != nil {
        count++
        c = c.Next
    }
    for i := 0 ; i < count/2 ; i++{
        head = head.Next
    }
    return head
}