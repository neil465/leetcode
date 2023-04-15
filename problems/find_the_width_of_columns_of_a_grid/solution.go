func findColumnWidth(grid [][]int) []int {
    res := []int{}
    for i := 0 ; i < len(grid[0]) ; i ++ {
        k := 0
        for _, j := range grid {
            a := strconv.Itoa(j[i])
            if k < len(a) {
                k = len(a)
            }
        }
        res = append(res, k)
        
    }
    return res
    
}