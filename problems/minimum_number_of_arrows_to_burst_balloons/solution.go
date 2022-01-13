func findMinArrowShots(points [][]int) int {
    sort.Slice(points,func(i,j int) bool {
        return points[i][1] > points[j][1]
    })
    sum := 1
    compare := points[0]
    for _,i := range points{
        left := max(compare[0],i[0])
        right := min(compare[1],i[1])
        if left > right{
            compare = i
            sum ++
        }else{
            compare = []int{left,right}
        }

        
    }
    return sum
}
func max(i,j int) int{
    if i > j {
        return i 
    }
    return j 
}
func min(i,j int) int{
    if i < j {
        return i 
    }
    return j 
}