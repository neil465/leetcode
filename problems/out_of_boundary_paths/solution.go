var dp = [50][50][51]int{}

func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
    for i := 0 ; i < 50 ; i++ { for j := 0 ; j < 50 ; j++ { for k := 0 ; k < 51 ; k++ { dp[i][j][k] = -1 } } }
    return helper(m,n,maxMove,startRow,startColumn) 
    
}

func helper(m, n, maxMove, startRow, startColumn int) int{
    
    if maxMove < 0 {
        return 0
    }
    if startRow < 0 || startRow >= m || startColumn < 0 || startColumn >= n {    
        return 1
    }
    if dp[startRow][startColumn][maxMove] != -1 {
        return dp[startRow][startColumn][maxMove]
    }
    
    dp[startRow][startColumn][maxMove] = ((((helper(m, n, maxMove - 1, startRow, startColumn - 1) % 1000000007) + helper(m, n, maxMove - 1, startRow, startColumn + 1)) % 1000000007 + helper(m, n, maxMove - 1, startRow - 1, startColumn)) % 1000000007 + helper(m, n, maxMove - 1, startRow + 1, startColumn)) % 1000000007
    return dp[startRow][startColumn][maxMove] 
}