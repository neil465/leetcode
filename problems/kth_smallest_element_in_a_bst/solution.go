/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
    if root == nil {
        return 0
    }
    stack := []*TreeNode{root}
   
    for root != nil || len(stack) > 0 {
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }  
        root = stack[len(stack) - 1]
        stack = stack[:len(stack) - 1]
        k-- //k-- to find the kth smallest element
        if k == 0 {
            return root.Val 
        }
        root = root.Right
    } 
    return 0
}