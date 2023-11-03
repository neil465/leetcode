/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type key struct {
    val *TreeNode
    r bool
}
func rob(root *TreeNode) int {
    dp := map[key] int {}
    return max(r(root, false, dp), r(root, true, dp))
}
func r(cur *TreeNode, shouldRob bool, dp map[key]int) int {
    if cur == nil {
        return 0
    }
    if val, ok := dp[key{cur, shouldRob}] ; ok {
        return val
    }
    k := 0
    if shouldRob {
        k = r(cur.Left, false, dp) + r(cur.Right, false, dp) + cur.Val
    } else {
        k = max(r(cur.Left, true, dp), r(cur.Left, false, dp)) + max(r(cur.Right, true, dp),r(cur.Right, false, dp))
    }
    dp[key{cur, shouldRob}] = k
    return k
}