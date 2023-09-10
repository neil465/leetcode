func minimumMoves(grid [][]int) int {

    zeros := [][]int{}
    for r := range grid {
        for c := range grid[0] {
            if grid[r][c] == 0{
                zeros = append(zeros, []int{r, c, 0})
            }
        }
    }
    return helper(grid, zeros, 0)
}
func helper(grid [][]int, zeros [][]int, depth int) int {
    b := false
    min := 10000000
    for z := range zeros {
        if zeros[z][2] == 0 {
            b  = true
            for nr := range grid {
                for nc := range grid[0] {
                    if grid[nr][nc] > 1{
                        grid[nr][nc] --
                        zeros[z][2] = 1
                        k := helper(grid, zeros, depth + abs(nr - zeros[z][0]) + abs(nc - zeros[z][1]))
                        grid[nr][nc] ++
                        zeros[z][2] = 0
                        if k < min {
                            min = k
                        }

                    }   
                }
            }
        }
    }
    if !b {
        return depth
    }
    return min
}
func abs(i int) int {
    if i < 0 {
        return -i
    }
    return i
}