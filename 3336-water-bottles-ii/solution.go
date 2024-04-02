var dp = map[[3]int]int{}

func maxBottlesDrunk(numBottles int, numExchange int) int {
    dp = map[[3]int]int{}
    return re(numBottles, numExchange, 0)
}
func re(bottles, exchange, empty int) int {
    if val, ok := dp[[3]int{bottles, exchange, empty}]; ok {
        return val
    }
    best := 0
    if bottles > 0 {
        best = re(bottles -1 , exchange, empty + 1) + 1
    }
    if empty >= exchange {
        best = max(best, re(bottles + 1, exchange + 1, empty - exchange))
    }
    dp[[3]int{bottles, exchange, empty}] = best
    return best
}
