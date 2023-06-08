func canAttendMeetings(intervals [][]int) bool {
    sort.Slice(intervals, func(i, j int) bool { 
        return intervals[i][0] < intervals[j][0] 
    })
    lend := -1 
    for i := 0; i < len(intervals); i++ {
        if intervals[i][0] < lend {
            return false
        }
        lend = intervals[i][1]
    }
    return true
}