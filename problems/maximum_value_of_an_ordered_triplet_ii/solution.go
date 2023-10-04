func maximumTripletValue(nums []int) int64 {
    pre, suf := make([]int, len(nums)), make([]int, len(nums))

    for i := 1 ; i < len(nums); i++{
        pre[i] = max(pre[i - 1], nums[i - 1])
    }
    for i := len(nums) - 2; i >= 0 ; i-- {
        suf[i] = max(suf[i + 1], nums[i + 1])
    }
    fmt.Println(pre, suf)
    max := 0
    for i := range nums {
        if ((pre[i] - nums[i]) * suf[i]) > max {
            max = ((pre[i] - nums[i]) * suf[i])
        }
    }
    return int64(max)
}
func max(i, j int) int {
    if i > j {
        return i 
    } else {
        return j
    }
}