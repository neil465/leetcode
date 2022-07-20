func checkXMatrix(grid [][]int) bool {
    isDiag := map[[2]int]bool{}
    
    for i,j :=0,0 ; i < len(grid) ; i,j = i+1,j+1 {
        if grid[i][j] == 0 {
            return false
        }
        isDiag[[2]int{i,j}] = true
    }
    for i,j :=0,len(grid) -1 ; i < len(grid) ; i,j = i+1,j-1 {
        if grid[i][j] == 0 {
            return false
        }
        isDiag[[2]int{i,j}] = true
    }
    for i := range grid {
        for j := range grid {
            if !isDiag[[2]int{i,j}] {
                if grid[i][j] != 0 {
                    return false
                }
            }
        }
    }
    return true
}