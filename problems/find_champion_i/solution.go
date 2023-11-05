func findChampion(grid [][]int) int {
    for i := range grid {
        k := false
        for j := range grid {
            
            if i != j && grid[i][j] == 0 {
                k = true
            }
            
        }
        if !k {
                return i
            }
    }
    return -1
}