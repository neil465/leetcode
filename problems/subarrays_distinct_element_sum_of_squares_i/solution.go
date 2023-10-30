func sumCounts(nums []int) int {
    v := 0 
    for i := 0 ; i < len(nums); i ++ {
        for j := i ; j < len(nums); j ++ {
            m := map[int]bool{}
            for k := i; k <= j; k ++ {
                m[nums[k]] = true
            }
            v += len(m) * len(m)
        }
    }
    return v
}