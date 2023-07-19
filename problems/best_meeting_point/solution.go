func minTotalDistance(grid [][]int) int {
    friends := [][]int{}
    for i := range grid {
        for j := range grid[i] {
            if grid[i][j] == 1 {
                friends = append(friends, []int{i, j})
            }
        }
    }
    min := math.MaxInt32
    for i := range grid {
        for j := range grid[i] {
            d := 0
            for _, k := range friends {
                d += abs(k[0] - i) + abs(k[1] - j)   
            }
            if d < min {
                min = d
            }
        }
    }
    return min
}
func abs(i int) int {
    if i < 0 {
        return -i
    }
    return i
}