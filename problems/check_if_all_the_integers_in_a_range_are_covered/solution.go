func isCovered(ranges [][]int, left int, right int) bool {
    m := make(map[int]bool)
    for _,i := range ranges{
        for j := i[0] ; j <= i[1] ;j++{
            m[j] = true
        }
    }
    for i := left ; i <=right;i++{
        if m[i] != true{return false}
    }
    return true
}