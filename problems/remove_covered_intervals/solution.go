func removeCoveredIntervals(intervals [][]int) int {
    sort.Slice(intervals, func(i, j int) bool {
        if intervals[i][1] == intervals[j][1] {
            return intervals[i][0] > intervals[j][0]
        }
        return intervals[i][1] < intervals[j][1]
    })

    for i := 0 ; i < len(intervals) - 1 ; i++ {
        if intervals[i][0] >= intervals[i + 1][0] && intervals[i][1] <= intervals[i + 1][1] {
            intervals = append(intervals[:i], intervals[i + 1:]...)
            i = -1
        }
    }
    return len(intervals)
}