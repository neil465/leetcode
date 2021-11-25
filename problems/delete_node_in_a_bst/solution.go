/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deleteNode(root *TreeNode, key int) *TreeNode {
    if root == nil { return nil }
    if key < root.Val{ 
        root.Left = deleteNode(root.Left,key) 
    }else if key > root.Val{ 
        root.Right = deleteNode(root.Right,key) 
    }else{
        if root.Left == nil && root.Right == nil{
            return nil
        }
        if root.Left != nil || root.Right != nil{
            if root.Left == nil { return root.Right}
            if root.Right == nil{ return root.Left }
        }
        minright := root.Right
        if minright.Left != nil{
            for minright.Left != nil{
                fmt.Println(minright)
                minright = minright.Left
            }
        }
        fmt.Println(root)
        root.Val = minright.Val
        root.Right = deleteNode(root.Right,minright.Val)
    }
    return root
    
}