/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type Solution struct {
    n []int
    l int
}


/** @param head The linked list's head.
        Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor(head *ListNode) Solution {
    num :=0
    g := []int{}
    curr := head
    for curr != nil {
        g = append(g,curr.Val)
        num++
        curr = curr.Next
    }
    return Solution{n:g,l:num}
}


/** Returns a random node's value. */
func (this *Solution) GetRandom() int {
    
    return this.n[rand.Intn(this.l)]
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */