func twoCitySchedCost(costs [][]int) int {
    sort.Slice(costs, func(i, j int) bool {
        
        return float64(costs[i][0])-float64(costs[i][1]) < float64(costs[j][0])-float64(costs[j][1])
    })

    sum := 0
    for i := 0 ; i < len(costs)/ 2; i++ {
        sum += costs[i][0]
    }

    for i := len(costs)/2 ; i < len(costs); i++ {
        sum += costs[i][1]
    }
    return sum
}