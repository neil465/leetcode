
func minCost(sPos []int, hPos []int, rCosts []int, cCosts []int) int {
    sum := 0
    for sPos[0] != hPos[0] || sPos[1] != hPos[1]  {
        if sPos[0] < hPos[0] {
            sPos[0] ++
            sum += rCosts[sPos[0]]
            
        }
        if sPos[0] > hPos[0] {
            sPos[0] --
            sum += rCosts[sPos[0]]
        }
        if sPos[1] < hPos[1] {
            sPos[1] ++
            sum += cCosts[sPos[1]]
        }
        if sPos[1] > hPos[1] {
            sPos[1] --
            sum += cCosts[sPos[1]]
        }
    }
    return sum
}
func max (i, j int) int {
    if i > j {
        return i
    }
    return j 
}

func min  (i, j int) int {
    if i < j {
        return i
    }
    return j 
}