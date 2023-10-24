/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var maxv = math.MinInt32

func maxPathSum(root *TreeNode) int {
  maxv = math.MinInt32 
  re(root)

  return maxv
}
func re(root *TreeNode) int {
  if root == nil {
    return 0
  }
  v1 , v2 := re(root.Right), re(root.Left)
  maxv = max(maxv, v1 + v2 + root.Val, v1 + root.Val, v2 + root.Val, root.Val)
  return max(  v1 + root.Val, v2 + root.Val, root.Val)

}
func max(a ...int) int {
  m := math.MinInt32
  for _, v := range a {
    if v > m {
      m = v
    }
  }
  return m
}