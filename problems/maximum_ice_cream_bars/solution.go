func maxIceCream(costs []int, coins int) int {
    sort.Ints(costs)
    res := 0
    for _, i := range costs{
        if i>coins{
            return res
        }else{
            res++
            coins-=i
        }
    }
    return res
}