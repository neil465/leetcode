func removeOnes(grid [][]int) bool {
    for _,i := range grid{
        if i[0] == 1 {
            for k := range i {
                i[k] = 1-i[k]
            }
        }
    }
    for j,i := range grid[0]{
        if i == 1 {
            for k := range grid {
                grid[k][j] = 1-grid[k][j]
            }
        }
    }
    for _,i := range grid{
        for _,j := range i{
            if j == 1{return false}
        }
    }
    return true
}