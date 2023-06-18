func findValueOfPartition(nums []int) int {
    sort.Ints(nums)
    min := 1010100100
    for i := 1 ; i < len(nums); i++ {
        if nums[i] - nums[i - 1] < min {
            min = nums[i] - nums[i - 1]
        }
    }
    if min ==1010100100{
        return 0
    }
    return min
}