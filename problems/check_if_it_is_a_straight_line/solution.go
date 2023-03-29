func checkStraightLine(coordinates [][]int) bool {
    last := 0.0
    for i := 1 ; i < len(coordinates); i++ {

        slope := float64(coordinates[i][1] - coordinates[i - 1][1]) / float64(coordinates[i][0] - coordinates[i - 1][0])
        if slope == math.Inf(-1) && last == math.Inf(1) {
            return true
        }

        if slope != last && last != 0.0{
            fmt.Println(slope, last)
            return false 
        }  
        last = slope
    }
    return true
}