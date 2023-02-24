func minimizeSum(nums []int) int {
    sort.Ints(nums)

    return min(nums[len(nums) - 1] - nums[2], nums[len(nums) - 3] - nums[0], nums[len(nums) - 2] - nums[1])
}
func min( i, j, k int) int {
    if i < j && i < k {
        return i
    } else if j < i && j < k {
        return j
    }
    return k
}