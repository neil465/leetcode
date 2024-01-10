/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var m = map[int][]int{}

func amountOfTime(root *TreeNode, start int) int {
    m = map[int][]int{}
    a(root)

    visited := map[int]bool{}

    queue := [][]int{[]int{start, 0}}

    max := 0

    for len(queue) > 0 {
        pop := queue[0]
        queue = queue[1:]

        if visited[pop[0]] {
            continue
        }

        visited[pop[0]] = true

        if pop[1] > max {
            max = pop[1]
        }
        for _, val := range m[pop[0]] {
            if visited[val] {
                continue
            }
            queue = append(queue, []int{val, pop[1] + 1})
        }
    }
    return max

}
func a(root *TreeNode)  {
    if root == nil {
        return
    }
    if root.Left != nil {
        m[root.Val] = append(m[root.Val], root.Left.Val)
        m[root.Left.Val] = append(m[root.Left.Val], root.Val)
    }
    if root.Right != nil {
        m[root.Val] = append(m[root.Val], root.Right.Val)
        m[root.Right.Val] = append(m[root.Right.Val], root.Val)
    }
    a(root.Left)    
    a(root.Right)
    
}