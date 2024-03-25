func minOperations(k int) int {

    best := math.MaxInt32
    for i := 1; i <= int(math.Ceil(float64(k)/float64(2))); i ++ {
        best = min(int(math.Ceil(float64(k)/float64(i))) + i - 2 , best)
    }
    return best
}
