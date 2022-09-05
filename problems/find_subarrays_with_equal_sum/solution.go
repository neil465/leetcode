func findSubarrays(nums []int) bool {
    m := map[int]bool{}
    for i := 0 ; i < len(nums) -1 ; i++{ 
        if m[ nums[i] + nums[i+1] ] {
            return true
        }
        m[nums[i]+nums[i+1]] = true
    }
    return false
}