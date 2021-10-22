/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 


type Codec struct {
    m  map[*TreeNode]*Node
    n int
}

func Constructor() *Codec {
    m := make(map[*TreeNode]*Node)
    return &Codec{m,0}
}

func (this *Codec) encode(root *Node) *TreeNode {
    a:=&TreeNode{Val:this.n}
    this.n++
    this.m[a] = root
    return a
}

func (this *Codec) decode(root *TreeNode) *Node {
    return this.m[root] 
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * bst := obj.encode(root);
 * ans := obj.decode(bst);
 */