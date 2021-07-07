/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func plusOne(head *ListNode) *ListNode {
    
    flag := false
    arr := []*ListNode{}
    l1 := &ListNode{Val:1}
    curr := head
    for curr != nil{
        arr = append(arr,curr)
        curr = curr.Next
    }
    for i := len(arr)-1;i>=0 ; i--{
        if arr[i].Val < 9{
            arr[i].Val = arr[i].Val+1
            break
        }else if i == 0{
            arr[0].Val = 0
            l1.Next = head
            flag = true
            break
        }else{
            arr[i].Val = 0
        }
    }
    if flag{
        return l1
    }
    return head
}