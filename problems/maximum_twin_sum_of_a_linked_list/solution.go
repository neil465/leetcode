/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func pairSum(head *ListNode) int {
    maxSum := 0
    i := -1
    var stack []int
    var slow = head
    var fast = head
    
    for slow != nil {
        if fast != nil {
            stack = append(stack,slow.Val)
            fast = fast.Next.Next
            i++
        }else{
            var pop = stack[i]
            if pop + slow.Val > maxSum{
                maxSum = pop + slow.Val
            }     
            i --
        }
        slow = slow.Next
        
    }
    
    return maxSum
}