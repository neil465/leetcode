func maxPoints(points [][]int) int {
    max := 0
    for _,i := range points{
        m := map[float64]int{}
        for _,j := range points{
            if !reflect.DeepEqual(i,j){
                m[float64(i[1]-j[1])/float64(i[0]-j[0])]++
            }
        }
        for _,i := range m{
            if i > max{
                max = i
            }
        }
    }
    
    return max+1
}