func maxOperations(nums []int) int {
    if len(nums) < 2 {
        return 0
    }
    checkval := nums[0] + nums[1]
    for i := 2 ; i < len(nums) - 1 ; i += 2 {
        if nums[i] + nums[i + 1] != checkval{
            return i / 2
        }
    }
    return len(nums) / 2
}