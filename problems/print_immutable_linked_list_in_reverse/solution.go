/*   Below is the interface for ImmutableListNode, which is already defined for you.
 *
 *   type ImmutableListNode struct {
 *       
 *   }
 *
 *   func (this *ImmutableListNode) getNext() ImmutableListNode {
 *		// return the next node.
 *   }
 *
 *   func (this *ImmutableListNode) printValue() {
 *		// print the value of this node.
 *   }
 */

func printLinkedListInReverse(head ImmutableListNode) {
    curr := head
    arr := []ImmutableListNode{}
    for curr != nil{
        arr = append(arr,curr)
        curr = ImmutableListNode.getNext(curr)
    }
    for i := len(arr)-1 ; i>=0; i--{
        ImmutableListNode.printValue(arr[i])
    }
    
}