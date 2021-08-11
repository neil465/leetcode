func heightChecker(heights []int) int {
    sum := 0
    dst := make([]int, len(heights))
    copy(dst,heights)
    sort.Ints(dst)
    for i,j := range heights{
        if j != dst[i]{sum++}
    }
    return sum
    
}