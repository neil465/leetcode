var dp = [100000]int{}

func minCostClimbingStairs(cost []int) int {
    dp = [100000]int{}
    return int(math.Min(float64(minClimb(cost,len(cost)-1)),float64(minClimb(cost,len(cost)-2))))
}
func minClimb(cost []int, pointer int) int {
    if dp[pointer] >0 {
        return dp[pointer] 
    }
    if pointer < 0 {
        return 0
    }
    if pointer == 0 || pointer == 1 {
        return cost[pointer]
    }
    dp[pointer] = cost[pointer] + int(math.Min(float64(minClimb(cost,pointer-1)),float64(minClimb(cost,pointer-2))))
    return dp[pointer]
}