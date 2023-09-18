func maxNumberOfAlloys(n int, k int, budget int, composition [][]int, stock []int, cost []int) int {

    start, end := 1, math.MaxInt32 
    max := 0
    for start <= end {
        
         mid := (start + end) / 2 
         f := true
        for ind := range composition {
           
            v :=  re( ind, mid, budget, composition, stock, cost)
            if v > -1 {
                start = mid + 1
                max = mid
                f = false
                break
            } 
        }     
        if f {
            end = mid - 1
        }

    }
    return max
    
}
func re( mind, alloys, budget int, composition [][]int, stock []int, cost []int) int { 
    t := make([]int, len(stock))
    copy(t, stock)

    for j := 0; j < len(composition[mind]); j++ {
        req := composition[mind][j] * alloys
        if req > t[j] {
            budget -= (req - t[j]) * cost[j]
        } 

        if budget < 0 {
            return -1
        }
    }

    return alloys
}
func max(i, j int) int {
    if i > j {
        return i
    }
    return j
}