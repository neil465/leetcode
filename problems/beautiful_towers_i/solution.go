func maximumSumOfHeights(maxHeights []int) int64 {
    m := 0
    for i := range maxHeights {
        sum := maxHeights[i]
        lastV := maxHeights[i]
        for j := i - 1 ; j >= 0 ; j -- {
            sum += min(maxHeights[j],lastV )
            lastV = min(maxHeights[j],lastV )
        }
        lastV = maxHeights[i]
        for j := i + 1 ; j < len(maxHeights); j ++ {
            sum += min(maxHeights[j], lastV)
            lastV = min(maxHeights[j],lastV )
        }
        if sum > m {

            m = sum
        }
    }
    return int64(m)
}
func min(i,j int)int {
    if i < j {
        return i
    }
    return j
}