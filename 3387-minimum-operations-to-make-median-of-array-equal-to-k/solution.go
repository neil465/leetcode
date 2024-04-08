func minOperationsToMakeMedianK(nums []int, k int) int64 {
    sort.Ints(nums)
    cur := len(nums) / 2
    ops := 0

    if nums[cur] < k {
        for cur = cur; cur < len(nums) && nums[cur] < k ; cur ++ {
            ops += k - nums[cur] 
        }        
    } else if nums[cur] > k{
        for cur = cur; cur >= 0 && nums[cur] > k; cur -- {
            ops += nums[cur] - k
        }
    } 
    return int64(ops)
}
