func shiftGrid(grid [][]int, k int) [][]int {
    
    var result = make([][]int,len(grid))
    var newI , newJ int 
    for i := 0 ; i < len(grid) ; i ++ {
        for j := 0 ; j < len(grid[i]) ; j++ {            
            newI , newJ = (i + (j + k) / len(grid[0])) % len(grid), (j + k) % len(grid[0])
            
            if len(result[newI]) == 0 {
                result[newI] = make([]int, len(grid[0]))
            }

            result[newI][newJ] = grid[i][j]
        }
    }
    return result
}