func uniquePaths(m int, n int) int {
    return pathFinder(m,n, map[string]int{})
}
func pathFinder(m int, n int, dp map[string]int) int {
    var pos = fmt.Sprint(m) + "," + fmt.Sprint(n)
    
    if (m == 1 || n == 1 ) || dp[pos] > 0 {
        return dp[pos] + 1 
    }
    
    dp[pos] = pathFinder(m - 1, n, dp) + pathFinder(m, n - 1, dp) - 1
    
    return dp[pos] + 1
}