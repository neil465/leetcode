func getRow(rowIndex int) []int {
    rowIndex ++ 
    start := [][]int{{1}}
    for i := 2 ; i <= rowIndex; i++ {
        next := make([]int, i)
        for j := range next {
            if j - 1 >= 0 {
                next[j] += start[len(start) -1][j - 1]
            }
            if j < len(start[len(start) -1]) {
                next[j] += start[len(start) -1][j]
            }
        }
        start = append(start, next)
    }
    return start[len(start) - 1]
}
