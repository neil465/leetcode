func isSameTree(r1 *TreeNode, r2 *TreeNode) bool {
    if r1 != nil && r2 != nil && r1.Val != r2.Val{return false}
    if r1 == nil && r2 == nil{return true}
    if (r1 == nil && r2 != nil)||(r1!=nil && r2==nil){return false}
	q := []*TreeNode{r1}
    q2 := []*TreeNode{r2}
    g,g2 := []int{},[]int{}
    helper(q,&g)
    helper(q2,&g2)
    if reflect.DeepEqual(g,g2){
        return true
    }
	return false
	
}
func helper(q []*TreeNode , g *[]int){
    if len(q)  == 0{return}
    pop := q[0]
    q = q[1:]
    if pop.Left != nil{
        *g = append(*g,pop.Left.Val)
        q = append(q,pop.Left)
    }
    if pop.Left == nil{
        *g = append(*g,-100002)
    }
    if pop.Right != nil{
        *g = append(*g,pop.Right.Val)
        q = append(q,pop.Right)
    }
    if pop.Right == nil{
        *g = append(*g,-100002)
    }
    helper(q,g)
}
