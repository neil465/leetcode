/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, low int, high int) int {
	res := 0
	stack := []*TreeNode{root}
	for len(stack)>0 {
		value := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
        if value != nil{
            if value.Val>=low && value.Val<=high {
			    res+= value.Val
		    }
		    if value.Val>low {
			    stack = append(stack,value.Left)
		    }
		    if value.Val<high {
			    stack = append(stack,value.Right)
		    }
        }
		
	}
	return res
}