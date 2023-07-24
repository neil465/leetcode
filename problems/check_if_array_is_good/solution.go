func isGood(nums []int) bool {
    if len(nums) == 0 {
        return true
    }
    sort.Ints(nums)
    if len(nums) != nums[len(nums) -1 ] + 1{
        return false
    }
    nums = nums[:len(nums) - 1]
    m := map[int]bool{}
    for i := 0; i < len(nums) - 1; i++ {
        
        if nums[i] + 1 != nums[i + 1] || m[nums[i]]{
            return false
        }
        m[nums[i]] = true
    }
    return true
}