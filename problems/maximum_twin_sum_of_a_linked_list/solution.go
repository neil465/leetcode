/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func pairSum(head *ListNode) int {
    i, stack ,maxSum, slow, fast := -1, []int{}, 0, head, head
    
    for slow != nil {
        if fast != nil { 
            // Until we reach the middle we are appending 
            // to the stack. 
            stack, fast, i = append(stack,slow.Val), fast.Next.Next, i + 1
        }else{
            // After we reach the middle checking the pop value + the 
            // nodes value to see if it is greater than the maxSum.
            var val = stack[i]
            if val + slow.Val > maxSum{
                maxSum = val + slow.Val
            }     
            i --
        }
        slow = slow.Next
        
    }
    
    return maxSum
}