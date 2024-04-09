var dp = map[int]int{}

func integerReplacement(n int) int {
    
    if n == 1 {
        return 0
    }
    if val, ok := dp[n]; ok {
        return val
    }
    if n % 2 == 0 {
        dp[n] = integerReplacement(n/2) + 1
    } else {
        dp[n] = min(integerReplacement(n - 1), integerReplacement(n + 1)) + 1
    }
    return dp[n]
}
