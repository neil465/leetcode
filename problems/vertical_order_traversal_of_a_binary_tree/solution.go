/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// same row sort

type tuple struct{
    x,y int
}
type F struct{
    r *TreeNode
    res int 
    y int
}
func verticalTraversal(root *TreeNode) [][]int {
    result :=  make(map[int][]tuple)
    r := F{root,0,0}
    helper(r,&result)
    fmt.Println(result)
    c := [][]int{}
    for i := -100000 ; i < 100000 ; i++{
        if len(result[i]) > 0{
            sort.Slice(result[i],func(k,j int )bool{
                if result[i][k].y == result[i][j].y{
                    return result[i][k].x < result[i][j].x
                }
                return result[i][k].y < result[i][j].y
            })
            arr := []int{}
            for _,b := range result[i]{
                arr = append(arr,b.x)
            }
            c = append(c,arr)
        }
        
    }
    return c
}
func helper(root F , result *map[int][]tuple) {
    b := *result
    b[root.res] = append(b[root.res],tuple{root.r.Val,root.y})
    *result = b
    if root.r.Right != nil{
        x := root.res +1 
        j := root.y+1
        helper(F{root.r.Right ,x,j } , result)
    }
    if root.r.Left != nil{
        x := root.res + -1
        j := root.y+1
        helper(F{root.r.Left ,x ,j} , result)
    }
    
}