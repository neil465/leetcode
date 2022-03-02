var dp []int

func climbStairs(n int) int {
    dp = make([]int, n + 1)
    return calClimb(n)
}
func calClimb(n int) int {
    if n < 0 { 
        return 0 
    }
    if n == 0 || dp[n] > 0  {
        return dp[n] + 1
    }
    dp[n] = calClimb(n - 1) + calClimb(n - 2) - 1
    return dp[n] + 1
}