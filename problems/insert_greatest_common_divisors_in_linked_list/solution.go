/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertGreatestCommonDivisors(head *ListNode) *ListNode {
    cur := head
    for cur.Next != nil {
        v := GCD(cur.Val, cur.Next.Val)
        k := cur.Next
        cur.Next = &ListNode{Val : v}
        cur.Next.Next = k
        cur = cur.Next.Next
    }
    return head
}
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}