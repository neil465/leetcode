func maximumStrongPairXor(nums []int) int {
    max := math.MinInt32
    for _, i := range nums {
        for _, j := range nums {
            if abs(i - j) <= min(i, j) && i ^ j > max {
                max = i ^ j 
            }
        }
    }
    return max
}
func abs(i int)  int {
    if i < 0 {
        return - i
    }
    return i
}