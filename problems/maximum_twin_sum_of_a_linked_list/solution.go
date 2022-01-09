/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func pairSum(head *ListNode) int {
    maxSum := 0
    var stack []int
    var slow = head
    var fast = head
    for slow != nil {
        if fast != nil {
            stack = append(stack,slow.Val)
            slow = slow.Next
            fast = fast.Next
            if fast != nil {
                fast = fast.Next
            }
        }else{
            if stack[len(stack)-1] + slow.Val > maxSum{
                maxSum = stack[len(stack)-1] + slow.Val
            } 
            slow = slow.Next 
            stack = stack[:len(stack)-1]
        }
        
    }
    return maxSum
    
}