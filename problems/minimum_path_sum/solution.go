var dp [1000][1000]int

func minPathSum(grid [][]int) int {
    dp = [1000][1000]int{}
    return helper(grid,len(grid)-1,len(grid[0])-1 )
}

func helper(grid [][]int,i,j int)int{
    
    if i < 0 || j < 0 {
        return 2147483647
    } else if i == 0 && j == 0 {
        return grid[i][j]
    }
    if dp[i][j] <= 0 { 
        dp[i][j] = grid[i][j] + min(helper(grid, i-1, j), helper(grid, i, j-1))
    }
    return dp[i][j]
}

func min (i, j int) int { if i < j { return i } ; return j }