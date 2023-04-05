func findLonelyPixel(picture [][]byte) int {
    rows, cols := map[int]int{}, map[int]int{}
    blackSpots := [][]int{}
    for i := range picture {
        for j := range picture[i] {
            if picture[i][j] == 'B' {
                blackSpots = append(blackSpots, []int{i, j})
                rows[i]++
                cols[j]++
            }
        }
    }
    count := 0
    fmt.Println(rows, cols)
    for _, i := range blackSpots {
        fmt.Println(rows[i[0]], cols[i[1]], i)
        if rows[i[0]] < 2 && cols[i[1]] < 2 {
            count++
        }
    }
    return count
}