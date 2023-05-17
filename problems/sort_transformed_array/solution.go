func sortTransformedArray(nums []int, a int, b int, c int) []int {
    for i := range nums {
        nums[i] = a * nums[i] * nums[i] + b * nums[i] + c
    }
    sort.Ints(nums)
    return nums
}