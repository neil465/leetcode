/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) int {
    res := 0
    curr := head
    arr := []int{}
    for curr != nil{
        arr = append(arr,curr.Val)
        curr = curr.Next
    }
    multiplyer := 1
    for i := len(arr)-1 ; i>=0 ; i--{
        res += multiplyer*arr[i]
        multiplyer*=2
    }    
    return res
}