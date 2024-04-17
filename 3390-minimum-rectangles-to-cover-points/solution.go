func minRectanglesToCoverPoints(points [][]int, w int) int {
    sort.Slice(points, func(i, j int) bool {
        return points[i][0] < points[j][0]
    })
    cur := 0 
    v := 0 
    for i := range points {
        if points[i][0] - points[cur][0] > w {
            v ++
            cur = i
        }
    }
    return v + 1
}
