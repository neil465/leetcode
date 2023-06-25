func numberOfGoodSubarraySplits(nums []int) int {
    lastindex := -1
    sum := 1
    b := false
    for i := range nums {
        if nums[i] == 1 {
            if lastindex == -1 {
                lastindex = i
                b = true
                continue
            }
            b = true

            sum *= (i - lastindex) 
            sum %=  1000000007
            lastindex = i
        }
    }
    if !b {
        return 0
    }
    return sum 
}