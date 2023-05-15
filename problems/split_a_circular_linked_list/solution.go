/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func splitCircularLinkedList(list *ListNode) []*ListNode {
    a := []*ListNode{}
    c := list
    for c != nil {
        a = append(a, c)

        c = c.Next
        if c == list {

            break
        }
    }
  
    a[int(math.Ceil(float64(len(a))/ float64(2))) - 1].Next = a[0]
    a[len(a) - 1].Next = a[int(math.Ceil(float64(len(a))/ float64(2)))]

    return []*ListNode{a[0], a[int(math.Ceil(float64(len(a))/ float64(2)))]}
}