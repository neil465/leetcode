/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type Node struct {
    node *TreeNode 
    depth int
}
func maxLevelSum(root *TreeNode) int {
    lv := map[int]int{}
    q := []Node{Node{root, 1}}
    minDepth := 0
    maxSum := math.MinInt32
    for len(q) > 0 {
        p := q[0]
        q = q[1:]
        if p.node == nil {
            continue
        }
        lv[p.depth] += p.node.Val
        q = append(q, Node{p.node.Left, p.depth + 1})
        q = append(q, Node{p.node.Right, p.depth + 1})
    }
    for k, v := range lv {
        if v > maxSum {
            minDepth = k
            maxSum = v
        }
    }
    return minDepth
}