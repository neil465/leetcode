func uniquePaths(m int, n int) int {
    return pathFinder(m,n, map[[2]int]int{})
}
func pathFinder(m int, n int, dp map[[2]int]int) int {
    var pos = [2]int{ m, n }
    
    if (m == 1 || n == 1 ) || dp[pos] > 0 {
        return dp[pos] + 1 
    }
    
    dp[pos] = pathFinder(m - 1, n, dp) + pathFinder(m, n - 1, dp) - 1
    
    return dp[pos] + 1
}