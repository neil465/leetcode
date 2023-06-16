func countServers(grid [][]int) int {
    rsum, csum, res := make([]int, len(grid)), make([]int, len(grid[0])), 0
    for i := range grid {for j := range grid[0] {if grid[i][j] == 1 {csum[j], rsum[i] = csum[j] + 1, rsum[i] + 1}}}
    for i := range grid {for j := range grid[0] {if (rsum[i] > 1 || csum[j] > 1) && grid[i][j] == 1 {res ++}}}
    return res
}