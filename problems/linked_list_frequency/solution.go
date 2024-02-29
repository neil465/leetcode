/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func frequenciesOfElements(head *ListNode) *ListNode {
    v := map[int]int{}
    cur := head
    for cur != nil {
        v[cur.Val] ++
        cur = cur.Next
    }

    resHead := &ListNode{}

    rescur := resHead 

    for _, val := range v {
        rescur.Next = &ListNode{}
        rescur.Next.Val = val
        
        rescur = rescur.Next
    }

    return resHead.Next
}