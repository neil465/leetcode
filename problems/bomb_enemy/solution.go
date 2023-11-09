func maxKilledEnemies(grid [][]byte) int {
    m := 0
    for i := range grid {
        for j := range grid[i] {
            if grid[i][j] == byte('0') {
                m = max(calc(grid, i, j), m)
                fmt.Println(i, j, m)
            }
        }
    }
    return m
}
func calc(grid [][]byte, r, c int) int {
    s := 0 
    for _, dir := range [][]int{{0,1}, {0,-1}, {1,0}, {-1,0}} {
        for i := 1 ; i <  max(len(grid), len(grid[0])); i ++ {
            if dir[0] * i + r >= 0 && dir[0] * i + r < len(grid) && dir[1] * i  + c>= 0 && dir[1] * i + c < len(grid[0]) && grid[r + dir[0] * i][c + dir[1] * i] != byte('W')  {
                if grid[r + dir[0] * i][c + dir[1] * i] == byte('E') {
                    s++
                }
            } else {
                break
            }
        }
    }
    return s
   
}