func equalPairs(grid [][]int) int {
    m := map[string]int{}
    a := 0
    for _, i := range grid {
        s := ""
        for _, j := range i {
            s += " " + fmt.Sprint(j) 
        }
        m[s]++

    }
    for i := range grid[0] {
        s := ""
        for j := range grid {
            s += " " + fmt.Sprint(grid[j][i])
        }
        a += m[s]
    }
    return a
}