/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
    if depth == 1 {
        a := &TreeNode{Val:val, Left: root}
        return a 
    }
    depth -= 2
    lastRow := []*TreeNode{root}

    for {
        if depth == 0 {
            for _, i := range lastRow {
                if i != nil {
                    left := &TreeNode{Val:val}
                    if i.Left != nil {
                        left.Left = i.Left
                    }
                    i.Left = left
                    

                    right := &TreeNode{Val:val}
                    if i.Right != nil {
                        right.Right = i.Right
                    }
                    i.Right = right
                }
            }
            break
        }
        depth --
        tempRow := []*TreeNode{}
        for _, i := range lastRow {
            if i.Left != nil {
                tempRow = append(tempRow, i.Left)
            }
            if i.Right != nil {
                tempRow = append(tempRow, i.Right)
            }
            
        }
        lastRow = tempRow
    }
    return root
    
}
