func minimumTime(jobs []int, workers []int) int {
    sort.Ints(jobs)
    sort.Ints(workers)
    max := 1
    for i := 0 ; i < len(jobs); i++{
        if temp := int(math.Ceil(float64(jobs[i])/float64(workers[i]))); temp > max {
            max = temp
        }
    }
    return max
}