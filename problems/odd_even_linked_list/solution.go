/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
    c := head
    odds,evens := &ListNode{},&ListNode{}
    curr,curr2 := odds,evens
    count := 0
    for c != nil{
        if count %2 != 0{
            l1 := &ListNode{Val:c.Val}
            curr2.Next = l1
            curr2 = curr2.Next
        }
        if count %2 == 0{
            l1 := &ListNode{Val:c.Val}
            curr.Next = l1
            curr = curr.Next
        }
        count ++
        c = c.Next
    }
    odds = odds.Next
    evens = evens.Next
    fmt.Println(odds,evens)
    curr.Next = evens
    return odds
}