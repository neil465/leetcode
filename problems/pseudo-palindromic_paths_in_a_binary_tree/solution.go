


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pseudoPalindromicPaths (root *TreeNode) int {
    return r(root, [10]int{})
}
func r(root *TreeNode, counts [10]int) int {
    
    if root == nil {
        return 0
    }

    counts[root.Val] ++
    if root.Left == nil && root.Right == nil {

        oddcount := 0 

        for _, count := range counts {
            if count % 2 != 0 {
                oddcount ++
            }
        }

        if oddcount > 1 {
            return 0
        }
        return 1
    }
    
    v := r(root.Left, counts) + r(root.Right , counts)
    return v
}