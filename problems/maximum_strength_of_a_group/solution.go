func maxStrength(nums []int) int64 {
    sort.Ints(nums)
    lNeg := 0
    s := 1
    nCount, pCount := 0, 0
    for  i := 0 ; i < len(nums); i ++ {
        if nums[i] == 0 {
            continue 
        }
        if nums[i] < 0 {
            lNeg = i
            nCount++
        } else {
            pCount++
        }
        s *= nums[i]
        
    }
    if s < 0 {
        s /= nums[lNeg]
        nCount--
    }
    if nCount == 0 && pCount == 0 {
        return int64(nums[len(nums) - 1])
    }
    return int64(s)
}