/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {
    a := []*Node{root}
    if root == nil {
        return [][]int{}
    }
    res := [][]int{{root.Val}}
    c := 0
    for len(a) > c {
        n := len(a)
        fmt.Println(res)
        res = append(res, []int{})
        for i := c ; i < n ; i ++{
            for _, v := range a[i].Children {
                res[len(res) - 1] = append(res[len(res) - 1], v.Val)
                a = append(a, v)
            }
        }
        c = n
    }
    return res[:len(res) -1]

}