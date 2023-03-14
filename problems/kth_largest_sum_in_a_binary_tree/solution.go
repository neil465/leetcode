/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthLargestLevelSum(root *TreeNode, k int) int64 {
    a := []int{root.Val}
    layer := []*TreeNode{root}
    for {
        if len(layer) == 0 {
            break
        }
        nextSum := 0
        nextLayer := []*TreeNode{}
        for _, i := range layer {
            if i.Left != nil {
                nextLayer = append(nextLayer, i.Left)
                nextSum += i.Left.Val
            }
            if i.Right != nil {
                nextLayer = append(nextLayer, i.Right)
                nextSum += i.Right.Val
            }
        }
        layer = nextLayer
        a = append(a, nextSum)
    }
    
    sort.Ints(a)
    if len(a) < k {
        return -1
    }
    if int64(a[len(a) - k ]) == 0 {
        return -1 
    }
    return int64(a[len(a) - k ])
}

