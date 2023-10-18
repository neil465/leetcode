func minCostII(costs [][]int) int {
    var dp = make([][]int, len(costs))
    for i := 0; i < len(costs) - 1; i ++ {
        dp[i] = make([]int, len(costs[i]))
    }    
    dp[len(dp) - 1] = costs[len(costs) - 1]
    for i := len(costs) -2 ; i >= 0 ; i -- {
        for j := range costs[0] {
            dp[i][j] = math.MaxInt32 
            for jk := range costs[0] {
                if dp[i + 1][jk] + costs[i][j] < dp[i][j] && jk != j {
                    dp[i][j] = dp[i + 1][jk] + costs[i][j]
                }
            }
        }   
    }
    min := math.MaxInt32
    for i := 0 ; i < len(dp[0]) ; i++ {
        if dp[0][i] < min {
            min = dp[0][i]
        }
    }
    return min
}
