func lexicalOrder(n int) []int {

    return find(n, 0)
}
func find(n int, cur int) []int {
    res := []int{}
    for i := 0; i < 10; i ++ {
        if i == 0 && cur == 0 {
            continue 
        }
        v := cur * 10 + i
        if v <= n {
            res = append(res, append([]int{v}, find(n, v)...)...)
        }
    }
    return res
}