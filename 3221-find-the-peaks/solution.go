func findPeaks(m []int) []int {
    inds := []int{}

    for i := 1; i < len(m) - 1; i ++ {
        if m[i] > m[i + 1] && m[i] > m[i - 1] {
            inds = append(inds, i)
        }
    }   
    return inds
}
