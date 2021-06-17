func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
    res := 0 
    for _,i := range arr1{
        flag := true
        for _,j := range arr2{
            if abs(i-j)<=d{
                flag = false
                break
            }
        }
        if flag {
            res++
        }
    }
    return res
}
func abs(i int) int{
    if i > 0 {
        return i
    }
    return i*-1
}