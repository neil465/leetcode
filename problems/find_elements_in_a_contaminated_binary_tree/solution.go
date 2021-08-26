/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type FindElements struct {
    a map[int]bool
}


func Constructor(root *TreeNode) FindElements {
    m := make(map[int]bool)
    root.Val = 0
    m[0] = true
    stack := []*TreeNode{root}
    for len(stack) > 0{
        pop := stack[len(stack)-1]
        
        stack = stack[:len(stack)-1]
        if pop.Right != nil{
            stack = append(stack,pop.Right)
            pop.Right.Val = (2*pop.Val)+2
            m[pop.Right.Val] = true
        }
        if pop.Left != nil{
            stack = append(stack,pop.Left)
            pop.Left.Val = (2*pop.Val)+1
            m[pop.Left.Val] = true
        }
        
    }
    return FindElements{m}
}


func (this *FindElements) Find(target int) bool {

    return this.a[target]
}


/**
 * Your FindElements object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Find(target);
 */