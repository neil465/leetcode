/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func replaceValueInTree(root *TreeNode) *TreeNode {
    a := []*TreeNode{root}
    root.Val = 0
    for len(a) > 0 {
        temp := []*TreeNode{}
        fmt.Println(a)
        m := map[int]bool{}
        sum := 0
        for _, i := range a {
            if i.Left != nil {
                temp = append(temp, i.Left)
                m[i.Left.Val] = true
                sum += i.Left.Val
            }
            if i.Right != nil {
                temp = append(temp, i.Right)
                m[i.Right.Val] = true 
                sum += i.Right.Val
            }
        }
        fmt.Println(sum)
        for _, i := range a {
            eResult := sum
            if i.Left != nil {
                eResult -=i.Left.Val
            }
            if i.Right != nil {
                eResult -=i.Right.Val
            }
            if i.Left != nil {
                i.Left.Val = eResult
            }
            if i.Right != nil {
                i.Right.Val = eResult
            }
        }
        a = temp
        
    }
    return root
    
}