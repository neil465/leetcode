func numberOfPoints(nums [][]int) int {
    m := map[int]bool{}
    for _, i := range nums {
        for k := i[0] ; k <= i[1] ; k ++ {
            m[k] = true
        }
    }
    return len(m)
}