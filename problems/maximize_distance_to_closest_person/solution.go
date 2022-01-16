func maxDistToClosest(seats []int) int {
    arr := []int{}
    for j,i := range seats{
        if i == 1{
            arr = append(arr,j)
        }
    }
    k := 0
    for i := 0 ; i < len(seats) ; i++{
        a := 1000000
        for _,j := range arr{
            if abs(i-j) < a {
                a = abs(i-j)
            }
        }
        if a > k{
            k = a
        }
    }
    return k
}
func abs(i int) int {
    if i > 0 {return i }
    return i*-1
}