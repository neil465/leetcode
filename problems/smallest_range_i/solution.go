func smallestRangeI(nums []int, k int) int {
    if len(nums) == 0 {
        return 0
    }
    min, max := nums[0], nums[0]
    for _, i := range nums {
        if i < min {
            min =i
        }
        if i >= max {
            max = i
        }
    } 
    if max - min - 2*k < 0 {
        return 0
    }
    return max - min - 2*k
}