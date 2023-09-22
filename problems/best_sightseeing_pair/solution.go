func maxScoreSightseeingPair(values []int) int {
    max := 0 
    current := 0 
    for _, i := range values {
        if current + i > max{ max = current + i}
        if i > current { current = i - 1} else { current --}
    }
    return max 
}