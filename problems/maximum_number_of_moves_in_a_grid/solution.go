var dp = map[[2]int]int{}

func maxMoves(grid [][]int) int {
    dp = map[[2]int]int{}
    m := 0
    for i := range grid {
        f := find(grid, i, 0)
        if f > m {
            m = f
        }
    }
    return m - 1
}
func find(grid [][]int, row, col int) int {
    if dp[[2]int{row, col}] > 0{
        return dp[[2]int{row, col}]
    }
    m := 0
    

    dirs := [][]int{{-1, 1}, {0, 1}, {1, 1}}
    for _, dir := range dirs {
        
        if row + dir[0] >= 0 && row + dir[0] < len(grid) && col + dir[1] >= 0 && col + dir[1] < len(grid[0]) && grid[row + dir[0]][col + dir[1]] > grid[row][col] {

            f := find(grid, row + dir[0], col + dir[1])

            if f > m {

                m = f
            }
        }
    }
    dp[[2]int{row, col}] = 1 +m 
    return 1 + m 
}