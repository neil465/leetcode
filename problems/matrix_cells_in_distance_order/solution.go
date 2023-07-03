func allCellsDistOrder(rows int, cols int, rCenter int, cCenter int) [][]int {
    b := [][]int{}
    for i := 0 ; i < rows; i ++ {
        for j := 0 ; j < cols; j++ {
            b = append(b, []int{i, j})
        }
    }
    sort.Slice(b, func(i, j int) bool {
        return abs(b[i][0] - rCenter) + abs(b[i][1] - cCenter) < abs(b[j][0] - rCenter) + abs(b[j][1] - cCenter)
    })
    return b
}
func abs(i int) int {
    if i < 0 {
        return i * - 1
    }
    return i
}