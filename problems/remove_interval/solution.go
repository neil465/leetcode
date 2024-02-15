func removeInterval(intervals [][]int, toBeRemoved []int) [][]int {
    arr := [][]int{}

    for _, interval := range intervals {
        if interval[1] < toBeRemoved[0] || interval[0] > toBeRemoved[1] {
            arr = append(arr, interval)
        } else if interval[0] > toBeRemoved[0] && interval[1] > toBeRemoved[1] {
            arr = append(arr, []int{toBeRemoved[1], interval[1]})
        } else if interval[0] < toBeRemoved[0] && interval[1] < toBeRemoved[1] {
            arr = append(arr, []int{interval[0], toBeRemoved[0]})
        } else if (interval[0] <= toBeRemoved[0] && interval[1] >= toBeRemoved[1]){
            if interval[0] != toBeRemoved[0] {
                arr = append(arr, []int{interval[0], toBeRemoved[0]})
            }
            if toBeRemoved[1] !=  interval[1] {
                arr = append(arr, []int{toBeRemoved[1], interval[1]})
            }
        }
    }
    return arr
}