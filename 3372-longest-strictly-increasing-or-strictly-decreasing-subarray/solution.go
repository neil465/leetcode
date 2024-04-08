func longestMonotonicSubarray(nums []int) int {
    best := 0
    for i := range nums {
        for j := i ; j < len(nums); j ++ {
            b := true
            for k := i + 1; k <= j; k ++ {
                if nums[k] <= nums[k - 1] {
                    b = false 
                    break
                }
            }
            if b {
                best = max(best, j - i + 1)   
            }
            b = true
            for k := i + 1; k <= j; k ++ {
                if nums[k] >= nums[k - 1] {
                    b = false 
                    break
                }
            }
            if b {
                best = max(best, j - i + 1)   
            }
        }
    }
    return best
}
