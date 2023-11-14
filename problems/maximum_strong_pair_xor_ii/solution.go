func maximumStrongPairXor(nums []int) int {
    sort.Ints(nums)
    m := 0 
    for i := 0 ; i < len(nums); i ++ {
        start, end := i, len(nums) - 1 
        mid := 0
        best := 0
        for start <= end {
            mid = (end - start) / 2 + start 
            if (nums[mid] - nums[i]) <= nums[i]  {
                best = mid
                start = mid + 1
            } else {
                end = mid - 1
            }
        }
        for indices := i ; indices <= best; indices ++ {
            m = max(m, nums[i] ^ nums[indices])
        }
    }
    return m
}