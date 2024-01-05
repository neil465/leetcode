type key struct {
    workerMask int
    bikeInd int
}

func assignBikes(workers [][]int, bikes [][]int) int {
    dp := map[key]int{}
    m := math.MaxInt32 

        m = min(m, r(dp, workers, bikes, 0, 0, 0))
    
    return m
}
func r(dp map[key]int, workers [][]int, bikes [][]int, bikeInd, mask, workersTaken int) int {
    
    best := math.MaxInt32 
    if len(workers) - workersTaken < len(bikes) - bikeInd  {
        best = r(dp, workers, bikes, bikeInd + 1, mask, workersTaken)
    }
    if val, ok := dp[key{mask, bikeInd}]; ok {
        return val
    }
    for i := range workers {
        if !isSet(i, mask) {
            best = min(best, r(dp, workers, bikes, bikeInd + 1, setBit(i, mask), workersTaken + 1) + dist(workers[i], bikes[bikeInd]))
        }
    }
    if best == math.MaxInt32 {
        return 0
    }
    dp[key{mask, bikeInd}] = best
    return best
}

func setBit(k, n int) int {
    return ((1 << k) | n)
}
func unsetBit(k, n int) int {
    return (^(1 << k) & n)
}
func isSet(k, n int) bool {
    return (1 << k) & n > 0
}
func dist(a, b []int) int {
    return int(math.Abs(float64(a[0] - b[0]))) + int(math.Abs(float64(a[1] - b[1])))
}