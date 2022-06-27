func minPartitions(n string) int {
    max := 0
    for i := 0 ; i < len(n); i++{
        max = int(math.Max(float64(max), float64(n[i] - '0')))
    }
    return max
}