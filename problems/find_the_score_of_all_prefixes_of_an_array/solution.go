func findPrefixScore(nums []int) []int64 {
    max := 0
    c := []int64{0}
    for _, i := range nums {
        if max < i {
            max = i
        }
        c = append(c, c[len(c) - 1] + int64(i+max))
        
    }
    return c[1:]
}