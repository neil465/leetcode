func findMissingAndRepeatedValues(grid [][]int) []int {
    m := map[int]int{}
    for i := range grid {
        for _, j := range grid[i] {
            m[j] ++
        }
    }
    dub, not := 0, 0
    for i := 1 ; i <= len(grid) * len(grid); i ++ {
        if m[i] < 1 {
            not = i
        } else if m[i] > 1 {
            dub = i
        }
    }
    return []int{dub, not}
}