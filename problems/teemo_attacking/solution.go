func findPoisonedDuration(timeSeries []int, duration int) int {
    res := 0
    for i := 0 ; i < len(timeSeries)-1 ; i++{
        if timeSeries[i] + duration-1 >= timeSeries[i+1]{
            res += timeSeries[i+1] - timeSeries[i]
        }else{
            res += duration
        }
    }
    res += duration
    return res
}