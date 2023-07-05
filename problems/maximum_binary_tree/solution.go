/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
    if len(nums) == 0 {
        return nil
    }
    b := 0
    for i := range nums {
        if nums[i] > nums[b] {
            b = i
        } 
    }
    node := &TreeNode{nums[b], constructMaximumBinaryTree(nums[:b]), constructMaximumBinaryTree(nums[b + 1:])}
    return node
}