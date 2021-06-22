/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    curr := head 
    arr := []int{}
    for curr != nil{
        if curr !=nil{
            arr = append(arr,curr.Val)
            curr = curr.Next
        }
        
        
    }
    if len(arr) == 0{
        return head
    }
    fmt.Println(arr)
    l := &ListNode{Val:arr[len(arr)-1]}
    curr2 := l  
    for i := len(arr)-2 ; i >= 0 ; i --{
        l1 := &ListNode{Val:arr[i]}
        curr2.Next = l1
        curr2 = curr2.Next
    }
    return l
}