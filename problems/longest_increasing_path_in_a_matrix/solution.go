var dp = [201][201]int{}

func longestIncreasingPath(matrix [][]int) int {
    dp = [201][201]int{}

    maxDepth := 0
    for i := range matrix {
        for j := range matrix[0] {
            r := re(i, j, 0, matrix)
            if r > maxDepth {
                maxDepth = r
            }
        }
    }
    return maxDepth + 1
}
func re(i, j, depth int, matrix [][]int) int {
    if dp[i][j] > 0{
        return dp[i][j] + depth
    }
    max := 0
    for _, dir := range [][]int{{0,1}, {0, -1}, {1,0}, {-1,0}} {
        fmt.Println(i + dir[0] >= 0 && i + dir[0] < len(matrix) && j + dir[1] >= 0 && j + dir[1] < len(matrix[0]), i + dir[0], j + dir[1], i, j, )
        if i + dir[0] >= 0 && i + dir[0] < len(matrix) && j + dir[1] >= 0 && j + dir[1] < len(matrix[0]) && matrix[i + dir[0]][j + dir[1]] > matrix[i][j] {
            r := re(i + dir[0], j + dir[1], 1, matrix)
            if r > max {
                max = r
            }
        }
    }
     dp[i][j] = max
    return depth + max
}