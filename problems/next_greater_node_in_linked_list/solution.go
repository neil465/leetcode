/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type key struct{
    k,v int
}
func nextLargerNodes(head *ListNode) []int {
    
    ms := []key{}
    m := make(map[int]int)
    
    count := 0 
    curr := head
    
    for curr != nil{
        for len(ms) > 0 && curr.Val > ms[len(ms)-1].k{
            m[ms[len(ms)-1:][0].v] = curr.Val
            ms = ms[:len(ms)-1]
        }
        ms = append(ms,key{k:curr.Val,v:count})
        
        curr = curr.Next
        count++
    }
    res := make([]int,count)
    for k,v := range m{
        res[k] = v
    }
    return res
}
