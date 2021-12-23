/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) int {
    s := ""
    for head != nil{
        s += fmt.Sprint(head.Val)
        head = head.Next
    }
    a, _ := strconv.ParseInt(s,2,64)
    return int(a)
}