/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
var res *ListNode
var cur *ListNode

func mergeKLists(lists []*ListNode) *ListNode {
    return helper(lists,res,cur)
}
func helper(lists []*ListNode , res *ListNode , cur *ListNode) *ListNode{
    removeIndex := -1
    for nodeIndex ,node := range lists{
        if removeIndex== -1 && node != nil{
            removeIndex = nodeIndex
        }else if node != nil && node.Val < lists[removeIndex].Val{
            removeIndex = nodeIndex
        }
    }
    if removeIndex == -1{
        return res
    }
    if res == nil {
        res = lists[removeIndex]
        cur = res
    }else{
        cur.Next = lists[removeIndex]
        cur = cur.Next
    }
    lists[removeIndex] = lists[removeIndex].Next
    helper(lists,res,cur)
    return res
}