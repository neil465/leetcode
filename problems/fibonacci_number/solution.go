var dp []int

func fib(n int) int {
    dp = make([]int, n+1)
    return calculate(n)
}
func calculate(n int) int {
    if n < 2 || dp[n] > 0 {
        if n == 0 {
            return 0
        }
        return dp[n] + 1
    }
    
    dp[n] = calculate(n-1) + calculate(n-2) - 1
    
    return dp[n] + 1
}