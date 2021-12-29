/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    arr :=[]*ListNode{}
    cur := head
    for cur != nil {
        arr = append(arr,cur)
        cur = cur.Next
    }
    if len(arr) - n == 0{
        head = head.Next
        return head
    }
    index := len(arr)-n
    fmt.Println(index-1)
    arr[index-1].Next = arr[index-1].Next.Next
   

    return head
}