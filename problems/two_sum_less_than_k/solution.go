func twoSumLessThanK(nums []int, k int) int {
    sort.Ints(nums)
    
    high, low, max := len(nums) - 1, 0, -1
    
    for low < high{
        if a := nums[high] + nums[low]; a  < k {
            if a > max {
                max = a
            }
            low++
        } else {
            high --
        }
    }
    return max
}