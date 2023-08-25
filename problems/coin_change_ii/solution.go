var dp = [5001][301]int{}

func change(amount int, coins []int) int {
    dp = [5001][301]int{}
    return find(0, amount, coins)
}
func find( i int, amount int, coins []int) int {
    if dp[amount][i] > 0 {
        return dp[amount][i] - 1
    }
    if amount == 0 {
        return 1
    } 
    g := 0
    for j := i; j < len(coins); j++ {
        if amount - coins[j] >= 0 {
            g += find(j, amount - coins[j], coins)
        }
    }
     dp[amount][i] = g + 1
    return dp[amount][i] - 1
}