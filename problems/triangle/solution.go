var dp [1000][1000]int
func minimumTotal(triangle [][]int) int {
    dp = [1000][1000]int{}
    for i := 999 ; i >= 0 ; i-- { 
        for j := 0 ; j < 1000 ; j++{
            dp[i][j] = 2147483647
        }
    }
    return pathFinder(0, 0, triangle)
}
func pathFinder(i, j int, triangle [][]int) int {
    if dp[i][j] != 2147483647 {
        return dp[i][j] 
    }
    if i == len(triangle) - 1 {
        return triangle[i][j]
    }
    dp[i][j] = triangle[i][j] + int(math.Min(float64(pathFinder(i+1, j, triangle)), float64(pathFinder(i+1, j+1, triangle)))) 
    return dp[i][j] 
}