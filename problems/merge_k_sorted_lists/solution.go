/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    if  len(lists) ==0 {
        return nil
    }
    for len(lists) > 1{
        a,b := lists[0],lists[1]
        lists =  lists[2:]
        merged := mergeTwoLists(a,b)
        
        t := []*ListNode{merged}
        
        lists = append(t, lists...)
    }
    return lists[0]
}
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoLists(l1, l2.Next)
	return l2
}