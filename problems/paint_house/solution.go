var dp = map[[3]int]int{}

func minCost(costs [][]int) int {
    dp = map[[3]int]int{}
    k := re(0, 0, costs, 0)
    return k
}
func re( curHouse, lastColor int, costs [][]int, sum int) (int) {

    if curHouse == len(costs) {
        return sum 
    }
    if _, ok := dp[[3]int{lastColor, curHouse, sum}] ; ok {
        return dp[[3]int{lastColor, curHouse, sum}]
    }
    min := math.MaxInt32
    for i := range costs[0] {
        if curHouse == 0 || i != lastColor {
            a := re(curHouse + 1, i, costs, sum + costs[curHouse][i])
            if a < min {
                min = a
            }
        }
    }
    dp[[3]int{lastColor, curHouse, sum}] = min
    return min

}