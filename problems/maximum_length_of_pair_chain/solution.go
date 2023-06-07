func findLongestChain(p [][]int) int {
    sort.Slice(p, func(i, j int) bool { 
        return p[i][1] < p[j][1] || (p[i][1] == p[j][1] && p[i][0] < p[j][0]) 
    })
    last, count := math.MinInt32, 0
    for _, pair := range p {
        if pair[0] > last {
            last = pair[1]
            count++
        }
    }
    return count
}
