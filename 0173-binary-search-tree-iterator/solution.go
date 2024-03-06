

import "sort"
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
    a []int
    cur int
}


func Constructor(root *TreeNode) BSTIterator {
    arr := re(root)
    sort.Ints(arr)
    return BSTIterator{arr, -1}
}
func re(i *TreeNode) []int {
    if i == nil {
        return []int{}
    }

    return append(re(i.Left), append(re(i.Right), i.Val)...)
}


func (this *BSTIterator) Next() int {
    this.cur = this.cur + 1
    return this.a[this.cur]
}


func (this *BSTIterator) HasNext() bool {
    return this.cur < len(this.a) - 1   
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
