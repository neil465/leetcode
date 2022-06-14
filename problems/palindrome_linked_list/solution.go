/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    s := []int{}
    slow := head
    for slow != nil {
        s = append(s, slow.Val)
        slow = slow.Next
    }
    fmt.Println(s)
    for i,j := 0, len(s) - 1; i <= j; i,j = i+1, j-1 {
        if s[i] != s[j] {
            return false
        }
    }
    return true
}